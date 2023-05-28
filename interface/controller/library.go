package controller

import "github.com/gin-gonic/gin"

type ILibraryController interface {
	AddNewBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
	FindBookById(c *gin.Context)
	FindAuthorById(c *gin.Context)
	FindAllBook(c *gin.Context)
}
