package handler

import (
	"net/http"

	"go-mongo/internal/model"
	"go-mongo/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) All(c *gin.Context) {
	users, err := h.service.All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Load(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.Load(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Insert(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	id := user.Id
	username := user.Username
	email := user.Email
	phone := user.Phone
	dateOfBirth := user.DateOfBirth

	newUser := model.User{
		Id:          id,
		Username:    username,
		Email:       email,
		Phone:       phone,
		DateOfBirth: dateOfBirth,
	}

	res, err := h.service.Insert(c.Request.Context(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	c.BindJSON(&user)
	err := h.service.Update(c.Request.Context(), &user, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted User!",
	})
}
