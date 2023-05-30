package main

import (
	"github.com/gin-gonic/gin"

	"fanchann/library/interface/controller"
	"fanchann/library/interface/middleware"
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
	utils.LogErrorWithPanic(err)

	bookRepo := books.NewBooksRepoImpl()
	bookInfosRepo := booksinformation.NewBooksInfoRepoImpl()
	authorRepo := authors.NewAuthorRepoImpl()
	services := services.NewLibraryImpl(db, authorRepo, bookRepo, bookInfosRepo)
	controller := controller.NewLibraryControllerImpl(services)

	router := gin.Default()
	router.Use(middleware.LibraryMiddleware())

	//Grouping
	group := router.Group("/libraries/")

	//For Author
	groupAuthor := group.Group("author/")
	//GET
	groupAuthor.GET(":id", controller.FindAuthorById)
	groupAuthor.GET("search", controller.FindAuthorByName)

	//For Book
	groupBook := group.Group("book/")
	//GET
	groupBook.GET(":id", controller.FindBookById)
	//POST
	groupBook.POST("new", controller.AddNewBook)
	//PUT
	groupBook.PUT(":id", controller.UpdateBook)
	//DELETE
	groupBook.DELETE(":id", controller.DeleteBook)

	//Get All Books And Authors
	//GET
	group.GET("books", controller.FindAllBook)
	group.GET("authors", controller.FindAllAuthorWithTheBook)

	router.Run("localhost:3000")
}
