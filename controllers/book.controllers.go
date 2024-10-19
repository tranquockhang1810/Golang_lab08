package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"lab07.com/database"
	"lab07.com/models"
)

//Get index
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Welcome to Book Management API!"})
}

// Get all books
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Create a new book
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// Update a book
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Delete a book
func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, book)
}

//Get a book
func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	database.DB.First(&book, id)
	c.JSON(http.StatusOK, book)
}

