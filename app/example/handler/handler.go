package handler

import (
	"go-gin/app/example"
	"go-gin/app/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ExampleHandler struct {
	uc  example.ExampleUsecase
	log *logrus.Entry
}

func ExampleRoute(uc example.ExampleUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := ExampleHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("example")

	v2.GET("", h.GetAllExample)
	v2.GET("/:id", h.GetDetailExample)
	v2.POST("", h.CreateExample)
	v2.PUT("/:id", h.UpdateExample)
	v2.DELETE("/:id", h.DeleteExample)
}

// @Tags Example
// @Summary Get All Example
// @Description Get All Example
// @Router /example [get]
// @Accept json
// @Produce json
// @param id query array false "ID in Array"
// @param name query string false "Filter by Name"
// @param username query string false "Filter by Username"
// @param email query string false "Filter by Email"
// @param limit query integer false "Limit of pagination"
// @param page query integer false "Page of pagination"
func (h *ExampleHandler) GetAllExample(c *gin.Context) {
	result, pagination, err := h.uc.GetAllExample(c)
	if err != nil {
		h.log.Errorf("get all example handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All Example",
		Meta:    pagination,
	})
}

// @Tags Example
// @Summary Get Detail Example
// @Description Get Detail Example by ID
// @Router /example/{id} [get]
// @Accept json
// @Produce json
// @param id path string true "ID"
func (h *ExampleHandler) GetDetailExample(c *gin.Context) {
	result, err := h.uc.GetDetailExample(c)
	if err != nil {
		h.log.Errorf("get detail example handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail Example",
	})
}

// @Tags Example
// @Summary Create Example
// @Description Create Example
// @Router /example [post]
// @Accept json
// @Produce json
// @Param request body example.ExampleForm true "Payload Body for Create [RAW]"
func (h *ExampleHandler) CreateExample(c *gin.Context) {
	err := h.uc.CreateExample(c)
	if err != nil {
		h.log.Errorf("create example handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "Create Example",
	})
}

// @Tags Example
// @Summary Update Example
// @Description Update Example
// @Router /example/{id} [put]
// @Accept json
// @Produce json
// @param id path string true "ID"
// @Param request body example.ExampleForm true "Payload Body for Update [RAW]"
func (h *ExampleHandler) UpdateExample(c *gin.Context) {
	err := h.uc.UpdateExample(c)
	if err != nil {
		h.log.Errorf("update example handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, tools.Response{
		Status:  "success",
		Message: "Update Example",
	})
}

// @Tags Example
// @Summary Delete Example
// @Description Delete Example by ID
// @Router /example/{id} [delete]
// @Accept json
// @Produce json
// @param id path string true "ID"
func (h *ExampleHandler) DeleteExample(c *gin.Context) {
	err := h.uc.DeleteExample(c)
	if err != nil {
		h.log.Errorf("delete example handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusNoContent, tools.Response{
		Status:  "success",
		Message: "Delete Example",
	})
}
