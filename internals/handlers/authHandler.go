package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/services"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/types/dto"
	"github.com/iacopoghilardi/golang-backend-boilerplate/pkg/validation"
	"github.com/iacopoghilardi/golang-backend-boilerplate/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (a *AuthHandler) Login(c *gin.Context) {
	var dto dto.LoginUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	token, err := a.authService.Login(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(token))
}

func (a *AuthHandler) Register(c *gin.Context) {
	var dto dto.RegisterUserDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	if err := validation.ValidateRegisterUserDto(&dto); err != nil {
		c.JSON(http.StatusBadRequest, utils.BuildErrorResponse(
			"Bad Request",
			err.Error(),
		))
		return
	}

	user, err := a.authService.Register(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse(
			"Internal Server Error",
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.BuildSuccessResponse(user))
}

func (a *AuthHandler) ResetPassword(c *gin.Context) {
}
