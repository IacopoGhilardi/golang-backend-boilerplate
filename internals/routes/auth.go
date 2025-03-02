package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/handlers"
)

func RegisterAuthRoutes(router *gin.RouterGroup, handler *handlers.AuthHandler) {
	router.POST("/login", handler.Login)
	router.POST("/register", handler.Register)
	router.POST("/reset-password", handler.ResetPassword)
}
