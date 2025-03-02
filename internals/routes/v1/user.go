package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/handlers"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/routes/middlewares"
)

func SetupUserRoutes(r *gin.RouterGroup, handler *handlers.UserHandler) {
	r.Use(middlewares.AuthMiddleware())

	r.GET("/", handler.GetAll)
	r.POST("/", handler.Create)
	r.GET("/:id", handler.GetById)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
