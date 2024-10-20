package controllers

import (
	"gin-gonic/auth"
	"gin-gonic/models"
	"gin-gonic/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	newUserReq := models.CreateUserRequest{}
	if err := c.ShouldBindJSON(&newUserReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUserID, err := uc.userService.CreateUser(&newUserReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.GenerateJWT(newUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	loginUserRequest := models.LoginUserRequest{}
	if err := c.ShouldBindJSON(&loginUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := uc.userService.LoginUser(&loginUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
