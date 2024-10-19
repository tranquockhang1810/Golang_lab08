package routes

import (
	"github.com/gin-gonic/gin"
	"lab07.com/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Định nghĩa các route cho sách
	router.GET("/", controllers.Index)
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books", controllers.UpdateBook)
	router.DELETE("/books", controllers.DeleteBook)
	return router
}