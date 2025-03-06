package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-gin/app/auth"
	authRepo "go-gin/app/auth/repository"
	authUC "go-gin/app/auth/usecase"
	mailerUC "go-gin/app/mailer/usecase"
	"go-gin/app/test"
	"go-gin/app/tools"
	"go-gin/app/user"
	userRepo "go-gin/app/user/repository"
	"go-gin/database/ent"
	entUser "go-gin/database/ent/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func ReturnAuthHandler(ctx context.Context, db *ent.Client, log *logrus.Entry) AuthHandler {
	repo := authRepo.NewAuthRepository(db)
	UserRepo := userRepo.NewUserRepository(db)
	mailerUC := &mailerUC.MockEmailUsecase{}
	ExampleUC := authUC.NewAuthUsecase(repo, UserRepo, ctx, mailerUC)

	return AuthHandler{
		uc:  ExampleUC,
		log: log,
	}
}

func TestRegister(t *testing.T) {
	// go test -v ./app/auth/handler/ -run TestRegister
	ht := test.SetUpRouter()

	h := ReturnAuthHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group("api/v1")
	v1.POST("/auth/register", h.Register)

	url := "/api/v1" + "/auth" + "/register"

	validData := user.UserForm{
		Name:     "example",
		Username: "example",
		Email:    "example@example.com",
		Password: "example",
	}

	invalidEmailandUsername := user.UserForm{
		Name:     "example testing",
		Username: "example",
		Email:    "example@example.com",
		Password: "example",
	}

	invalidEmail := user.UserForm{
		Name:     "example testing",
		Username: "example123",
		Email:    "example@example.com",
		Password: "example",
	}

	invalidUsername := user.UserForm{
		Name:     "example testing",
		Username: "example",
		Email:    "example123@example.com",
		Password: "example",
	}

	tests := []struct {
		name            string
		request         *user.UserForm
		expectedMessage interface{}
		expectedCode    int
		wantErr         bool
	}{
		// {"valid data", &validInputData, 201, false},
		{"valid register", &validData, "success register", 201, false},
		{"invalid email and username", &invalidEmailandUsername, auth.ErrEmailUsernameExists.Error(), 400, true},
		{"invalid email", &invalidEmail, auth.ErrEmailExists.Error(), 400, true},
		{"invalid username", &invalidUsername, auth.ErrUsernameExists.Error(), 400, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))

			w := httptest.NewRecorder()
			ht.Route.ServeHTTP(w, req)
			fmt.Println("result body", w.Body.String())
			fmt.Println("result code", w.Code)
			var response tools.Response
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Logf("gagal unmarshal respons: %v", err)
			}

			assert.Equal(t, tt.expectedCode, w.Code, "Code: Expected and Result does not match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message: Expected and Result does not match")
		})
	}

}

func TestActivatedClient(t *testing.T) {
	// go test -v ./app/auth/handler/ -run TestActivatedClient
	ht := test.SetUpRouter()

	h := ReturnAuthHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group("api/v1")
	v1.POST("/auth/activated-client", h.ActivatedClient)

	url := "/api/v1" + "/auth" + "/activated-client"

	userQuery := ht.DB.User.Query().
		Where(
			entUser.Or(
				entUser.UsernameEqualFold("example"),
			),
		)

	if count, _ := userQuery.Count(ht.Ctx); count == 0 {
		panic("not found user username")
	}

	exec, _ := userQuery.First(ht.Ctx)

	AT, err := h.uc.GetDetailAT(exec.ID)
	if err != nil {
		panic("not found user AT")
	}

	validData := auth.ActivatedTokenForm{
		Token: AT.Token,
	}

	tests := []struct {
		name            string
		request         *auth.ActivatedTokenForm
		expectedMessage interface{}
		expectedCode    int
		wantErr         bool
	}{
		{"account success activated", &validData, "success activated", 200, false},
		{"token already use", &validData, auth.ErrATalreadyUsed.Error(), 401, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tx, err := ht.DB.Tx(ctx) // Mulai transaksi
			if err != nil {
				t.Fatalf("gagal memulai transaksi: %v", err)
			}
			defer tx.Rollback() // Pastikan rollback selalu dijalankan

			// Kirim transaksi ke handler melalui context
			ctx = context.WithValue(ctx, "tx", tx)
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
			req = req.WithContext(ctx) // Set context ke request

			w := httptest.NewRecorder()
			ht.Route.ServeHTTP(w, req)

			var response tools.Response
			err = json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Logf("gagal unmarshal respons: %v", err)
			}

			// t.Log("result body", w.Body.String())
			// t.Log("result code", w.Code)
			assert.Equal(t, tt.expectedCode, w.Code, "Code: Expected and Result does not match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message: Expected and Result does not match")

			if !tt.wantErr {
				// dataBytes, err := json.Marshal(response.Data)
				// if err != nil {
				// 	t.Fatalf("gagal marshal response.Data: %v", err)
				// }
				// var data auth.ActivatedTokenForm
				// err = json.Unmarshal(dataBytes, &data)
				// if err != nil {
				// 	t.Logf("gagal unmarshal data: %v", err)
				// }
				// if data.Token == "" || data.RefreshToken == "" {
				// 	t.Error("not return token and refresh token")
				// }
			}

		})
	}

}

func TestLogin(t *testing.T) {
	// go test -v ./app/auth/handler/ -run TestLogin
	ht := test.SetUpRouter()

	h := ReturnAuthHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group("api/v1")
	v1.POST("/auth/login", h.Login)

	url := "/api/v1" + "/auth" + "/login"

	validData := auth.LoginForm{
		Username: "example",
		Email:    "example@example.com",
		Password: "example",
	}

	invalidEmailandUsername := auth.LoginForm{
		Username: "example12345",
		Email:    "example12345@example.com",
		Password: "example",
	}

	userNotActive := auth.LoginForm{
		Username: "existing",
		Email:    "exaexistingmple@example.com",
		Password: "existing",
	}

	invalidPassword := auth.LoginForm{
		Username: "example",
		Email:    "example@example.com",
		Password: "example1234",
	}

	tests := []struct {
		name            string
		request         *auth.LoginForm
		expectedMessage interface{}
		expectedCode    int
		wantErr         bool
	}{
		// {"valid data", &validInputData, 201, false},
		{"valid login", &validData, "success login", 200, false},
		{"account has not been activated", &userNotActive, auth.ErrAccountNotActivated.Error(), 403, true},
		{"invalid email and username", &invalidEmailandUsername, auth.ErrIdentityNotFound.Error(), 404, true},
		{"invalid password", &invalidPassword, auth.ErrWrongPassword.Error(), 401, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			tx, err := ht.DB.Tx(ctx) // Mulai transaksi
			if err != nil {
				t.Fatalf("gagal memulai transaksi: %v", err)
			}
			defer tx.Rollback() // Pastikan rollback selalu dijalankan

			// Kirim transaksi ke handler melalui context
			ctx = context.WithValue(ctx, "tx", tx)
			reqBody, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))
			req = req.WithContext(ctx) // Set context ke request

			w := httptest.NewRecorder()
			ht.Route.ServeHTTP(w, req)

			var response tools.Response
			err = json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Logf("gagal unmarshal respons: %v", err)
			}

			t.Log("result body", w.Body.String())
			t.Log("result code", w.Code)
			assert.Equal(t, tt.expectedCode, w.Code, "Code: Expected and Result does not match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message: Expected and Result does not match")

			if !tt.wantErr {
				dataBytes, err := json.Marshal(response.Data)
				if err != nil {
					t.Fatalf("gagal marshal response.Data: %v", err)
				}
				var data auth.Authenticate
				err = json.Unmarshal(dataBytes, &data)
				if err != nil {
					t.Logf("gagal unmarshal data: %v", err)
				}
				t.Log("data", data)
				if data.AccessToken.Token == "" || data.AccessToken.RefreshToken == "" {
					t.Error("not return token and refresh token")
				}
			}

		})
	}

}
