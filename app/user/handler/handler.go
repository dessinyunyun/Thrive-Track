package handler

import (
	"go-gin/app/tools"
	"go-gin/app/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	uc  user.UserUsecase
	log *logrus.Entry
}

func UserRoute(uc user.UserUsecase, r *gin.RouterGroup, log *logrus.Entry) {
	h := UserHandler{
		uc:  uc,
		log: log,
	}

	v2 := r.Group("user")

	v2.GET("", h.GetAllUser)
	v2.GET("/:id", h.GetDetailUser)
	v2.POST("", h.CreateUser)
	v2.PUT("/:id", h.UpdateUser)
	v2.DELETE("/:id", h.DeleteUser)
}

// @Tags User
// @Summary Get All Users
// @Description Get All Example
// @Router /user [get]
// @Accept json
// @Produce json
// @param id query array false "ID in Array"
// @param name query string false "Filter by Name"
// @param username query string false "Filter by Username"
// @param email query string false "Filter by Email"
// @param limit query integer false "Limit of pagination"
// @param page query integer false "Page of pagination"
func (h *UserHandler) GetAllUser(c *gin.Context) {
	result, pagination, err := h.uc.GetAllUser(c)
	if err != nil {
		h.log.Errorf("get all User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get All User",
		Meta:    pagination,
	})
}

// @Tags User
// @Summary Get Detail User
// @Description Get Detail User by ID
// @Router /user/{id} [get]
// @Accept json
// @Produce json
// @param id path string true "ID"
func (h *UserHandler) GetDetailUser(c *gin.Context) {
	result, err := h.uc.GetDetailUser(c)
	if err != nil {
		h.log.Errorf("get detail User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.Response{
		Data:    result,
		Status:  "success",
		Message: "Get Detail User",
	})
}

// @Tags User
// @Summary Create User
// @Description Create User
// @Router /user [post]
// @Accept json
// @Produce json
// @Param request body user.UserForm true "Payload Body for Create [RAW]"
func (h *UserHandler) CreateUser(c *gin.Context) {
	err := h.uc.CreateUser(c)
	if err != nil {
		h.log.Errorf("create User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, tools.Response{
		Status:  "success",
		Message: "Create User",
	})
}

// @Tags User
// @Summary Update User
// @Description Update User
// @Router /user/{id} [put]
// @Accept json
// @Produce json
// @param id path string true "ID"
// @Param request body user.UserForm true "Payload Body for Update [RAW]"
func (h *UserHandler) UpdateUser(c *gin.Context) {
	err := h.uc.UpdateUser(c)
	if err != nil {
		h.log.Errorf("update User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusAccepted, tools.Response{
		Status:  "success",
		Message: "Update User",
	})
}

// @Tags User
// @Summary Delete User
// @Description Delete User by ID
// @Router /user/{id} [delete]
// @Accept json
// @Produce json
// @param id path string true "ID"
func (h *UserHandler) DeleteUser(c *gin.Context) {
	err := h.uc.DeleteUser(c)
	if err != nil {
		h.log.Errorf("delete User handlers: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusNoContent, tools.Response{
		Status:  "success",
		Message: "Delete User",
	})
}
