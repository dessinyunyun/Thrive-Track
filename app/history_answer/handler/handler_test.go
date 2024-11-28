package handler

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"go-gin/database/ent"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/sirupsen/logrus"
// 	"gotest.tools/assert"

// 	// "go-gin/app/example"
// 	// "go-gin/app/example"
// 	"go-gin/app/test"
// 	"go-gin/app/user"
// 	userRepo "go-gin/app/user/repository"
// 	userUC "go-gin/app/user/usecase"
// )

// var ID uuid.UUID

// type Result struct {
// 	Data []user.FilterUser `json:"data"`
// }

// func ReturnUserHandler(ctx context.Context, db *ent.Client, log *logrus.Entry) Handler {
// 	UserRepo := userRepo.NewUserRepository(db)
// 	UserUC := userUC.NewUserUsecase(UserRepo, ctx)

// 	return Handler{
// 		uc:  UserUC,
// 		log: log,
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	ht := test.SetUpRouter()

// 	h := ReturnUserHandler(ht.Ctx, ht.DB, ht.Log)
// 	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
// 	v1.POST("User", h.CreateUser)

// 	body := user.UserForm{
// 		Name:     "UserRepo testing",
// 		Username: "UserRepo",
// 		Email:    "UserRepo@User.com",
// 	}

// 	url := os.Getenv("PREFIX_API") + "/" + "UserRepo"

// 	reqBody, _ := json.Marshal(body)

// 	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))

// 	w := httptest.NewRecorder()
// 	ht.Route.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// }

// func TestGetAllUser(t *testing.T) {
// 	ht := test.SetUpRouter()

// 	h := ReturnUserHandler(ht.Ctx, ht.DB, ht.Log)
// 	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
// 	v1.GET("user", h.GetAllUser)

// 	url := os.Getenv("PREFIX_API") + "/" + "user" + "?" + "limit=1"

// 	req, _ := http.NewRequest(http.MethodGet, url, nil)

// 	w := httptest.NewRecorder()
// 	ht.Route.ServeHTTP(w, req)

// 	body, _ := ioutil.ReadAll(w.Body)

// 	var result Result
// 	_ = json.Unmarshal(body, &result)

// 	ID = result.Data[0].ID

// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestGetDetailUser(t *testing.T) {
// 	ht := test.SetUpRouter()

// 	h := ReturnUserHandler(ht.Ctx, ht.DB, ht.Log)
// 	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
// 	v1.GET("User/:id", h.GetDetailUser)

// 	url := os.Getenv("PREFIX_API") + "/" + "User" + "/" + ID.String()

// 	req, _ := http.NewRequest(http.MethodGet, url, nil)

// 	w := httptest.NewRecorder()
// 	ht.Route.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// func TestUpdateUser(t *testing.T) {
// 	ht := test.SetUpRouter()

// 	h := ReturnUserHandler(ht.Ctx, ht.DB, ht.Log)
// 	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
// 	v1.PUT("User/:id", h.UpdateUser)

// 	body := user.UserForm{
// 		Name:     "User testing",
// 		Username: "User",
// 		Email:    "User@example.com",
// 	}

// 	reqBody, _ := json.Marshal(body)

// 	url := os.Getenv("PREFIX_API") + "/" + "User" + "/" + ID.String()

// 	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(reqBody))

// 	w := httptest.NewRecorder()
// 	ht.Route.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusAccepted, w.Code)
// }

// func TestDeleteUser(t *testing.T) {
// 	ht := test.SetUpRouter()

// 	h := ReturnUserHandler(ht.Ctx, ht.DB, ht.Log)
// 	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
// 	v1.DELETE("User/:id", h.DeleteUser)

// 	url := os.Getenv("PREFIX_API") + "/" + "User" + "/" + ID.String()

// 	req, _ := http.NewRequest(http.MethodDelete, url, nil)

// 	w := httptest.NewRecorder()
// 	ht.Route.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusNoContent, w.Code)
// }
