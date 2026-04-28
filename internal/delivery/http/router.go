package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irvanrifai/go-clean-architecture/internal/delivery/http/handler"
)

func NewRouter(handler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		v1.POST("/register", handler.Register)
		// v1.POST("/login", handler.Login)
	}

	return r
}
