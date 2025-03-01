package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"go-gin/app/auth"
	authUC "go-gin/app/auth/usecase"
	"go-gin/app/test"
	"go-gin/app/tools"
	"go-gin/app/user"
	userRepo "go-gin/app/user/repository"
	"go-gin/database/ent"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

func ReturnAuthHandler(ctx context.Context, db *ent.Client, log *logrus.Entry) AuthHandler {
	ExampleRepo := userRepo.NewUserRepository(db)
	ExampleUC := authUC.NewAuthUsecase(ExampleRepo, ctx)

	return AuthHandler{
		uc:  ExampleUC,
		log: log,
	}
}

func TestRegister(t *testing.T) {
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
		{"invalid email and username", &invalidEmailandUsername, auth.ErrEmailandUsernameAlreadyExist.Error(), 400, true},
		{"invalid email", &invalidEmail, auth.ErrEmailAlreadyExist.Error(), 400, true},
		{"invalid username", &invalidUsername, auth.ErrUsernameAlreadyExist.Error(), 400, true},
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
		})
	}

}

func TestLogin(t *testing.T) {
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

	validUsernameOnly := auth.LoginForm{
		Username: "example",
		Email:    "example12345@example.com",
		Password: "example",
	}

	validEmailOnly := auth.LoginForm{
		Username: "example12345",
		Email:    "example@example.com",
		Password: "example",
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
		{"invalid email and username", &invalidEmailandUsername, auth.ErrIdentityNotFound.Error(), 404, true},
		{"valid email only", &validEmailOnly, "success login", 200, false},
		{"valid username only", &validUsernameOnly, "success login", 200, false},
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

			// t.Log("result body", w.Body.String())
			// t.Log("result code", w.Code)
			assert.Equal(t, tt.expectedCode, w.Code, "Code: Expected and Result does not match")
			assert.Equal(t, tt.expectedMessage, response.Message, "message: Expected and Result does not match")

			if !tt.wantErr {
				dataBytes, err := json.Marshal(response.Data)
				if err != nil {
					t.Fatalf("gagal marshal response.Data: %v", err)
				}
				var data auth.LoginResponse
				err = json.Unmarshal(dataBytes, &data)
				if err != nil {
					t.Logf("gagal unmarshal data: %v", err)
				}
				if data.Token == "" || data.RefreshToken == "" {
					t.Error("not return token and refresh token")
				}
			}

		})
	}

}
