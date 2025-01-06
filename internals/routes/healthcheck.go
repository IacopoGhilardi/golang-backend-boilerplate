package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/db"
	"github.com/iacopoghilardi/golang-backend-boilerplate/utils"
)

func SetupHealthcheckRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("Database connection failed", err.Error()))
			return
		}

		c.JSON(http.StatusOK, utils.BuildSuccessResponse("OK"))
	})
}
