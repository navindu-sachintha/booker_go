package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/navindu-sachintha/go-bookstore/pkg/controllers"
)

func SetupRouter(handler *controllers.Handler) *gin.Engine {
	r := gin.Default()
	// Server test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/book", controllers.NewHandler().HandleCreateBook)
	r.GET("/books", controllers.NewHandler().HandleListBooks)
	r.GET("/book/:id", controllers.NewHandler().HandleGetBook)
	r.PUT("/book/:id", controllers.NewHandler().HanldeUpdateBook)
	r.DELETE("/deletebook/:id", controllers.NewHandler().DeleteBook)

	return r
}
