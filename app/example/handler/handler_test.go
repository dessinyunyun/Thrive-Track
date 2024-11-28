package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"go-gin/database/ent"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gotest.tools/assert"

	"go-gin/app/example"
	exampleRepo "go-gin/app/example/repository"
	exampleUC "go-gin/app/example/usecase"
	"go-gin/app/test"
)

var ID uuid.UUID

type Result struct {
	Data []example.FilterExample `json:"data"`
}

func ReturnExampleHandler(ctx context.Context, db *ent.Client, log *logrus.Entry) ExampleHandler {
	ExampleRepo := exampleRepo.NewExampleRepository(db)
	ExampleUC := exampleUC.NewExampleUsecase(ExampleRepo, ctx)

	return ExampleHandler{
		uc:  ExampleUC,
		log: log,
	}
}

func TestCreateExample(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnExampleHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.POST("example", h.CreateExample)

	body := example.ExampleForm{
		Name:     "example testing",
		Username: "example",
		Email:    "example@example.com",
	}

	url := os.Getenv("PREFIX_API") + "/" + "example"

	reqBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllExample(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnExampleHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.GET("example", h.GetAllExample)

	url := os.Getenv("PREFIX_API") + "/" + "example" + "?" + "limit=1"

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	var result Result
	_ = json.Unmarshal(body, &result)

	ID = result.Data[0].ID

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDetailExample(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnExampleHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.GET("example/:id", h.GetDetailExample)

	url := os.Getenv("PREFIX_API") + "/" + "example" + "/" + ID.String()

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateExample(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnExampleHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.PUT("example/:id", h.UpdateExample)

	body := example.ExampleForm{
		Name:     "example testing",
		Username: "example",
		Email:    "example@example.com",
	}

	reqBody, _ := json.Marshal(body)

	url := os.Getenv("PREFIX_API") + "/" + "example" + "/" + ID.String()

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(reqBody))

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
}

func TestDeleteExample(t *testing.T) {
	ht := test.SetUpRouter()

	h := ReturnExampleHandler(ht.Ctx, ht.DB, ht.Log)
	v1 := ht.Route.Group(os.Getenv("PREFIX_API"))
	v1.DELETE("example/:id", h.DeleteExample)

	url := os.Getenv("PREFIX_API") + "/" + "example" + "/" + ID.String()

	req, _ := http.NewRequest(http.MethodDelete, url, nil)

	w := httptest.NewRecorder()
	ht.Route.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
