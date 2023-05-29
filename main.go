package main

import (
	"github.com/gin-gonic/gin"

	"fanchann/library/interface/controller"
	"fanchann/library/internal/repositories/authors"
	"fanchann/library/internal/repositories/books"
	booksinformation "fanchann/library/internal/repositories/books_information"
	"fanchann/library/internal/services"
	"fanchann/library/pkg/database"
	"fanchann/library/pkg/utils"
)

func main() {
	db, err := database.MysqlConnect()
	utils.LogErrorWithPanic(err)
	err = db.Ping()
	utils.LogErrorWithPanic(err)

	bookRepo := books.NewBooksRepoImpl()
	bookInfosRepo := booksinformation.NewBooksInfoRepoImpl()
	authorRepo := authors.NewAuthorRepoImpl()
	services := services.NewLibraryImpl(db, authorRepo, bookRepo, bookInfosRepo)
	controller := controller.NewLibraryControllerImpl(services)

	// //success
	// fmt.Println("Success connected to database")

	router := gin.Default()

	router.GET("/libraries/books", controller.FindAllBook)
	router.GET("/libraries/book/:id", controller.FindBookById)
	router.GET("/libraries/authors", controller.FindAuthorById)
	router.GET("/libraries/author/:id", controller.FindAuthorById)
	router.GET("/libraries/author/search", controller.FindAuthorByName)

	router.POST("/libraries/books/new", controller.AddNewBook)
	router.PUT("/libraries/books/:id", controller.UpdateBook)
	router.DELETE("/libraries/books/:id", controller.DeleteBook)

	router.Run("localhost:3000")
}
