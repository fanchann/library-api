package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"fanchann/library/internal/models/web"
	"fanchann/library/internal/services"
	"fanchann/library/pkg/utils"
)

type LibraryControllerImpl struct {
	Service services.ILibraryServices
}

func NewLibraryControllerImpl(service services.ILibraryServices) ILibraryController {
	return &LibraryControllerImpl{Service: service}
}

func (cntrl *LibraryControllerImpl) AddNewBook(c *gin.Context) {
	form := web.AddNewBooks{}
	if errBind := c.BindJSON(&form); errBind != nil {
		utils.LogErrorWithPanic(errBind)
	}

	if form.Author == "" || form.Book_Title == "" {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "failed add new book", nil))
		return
	}

	dataForm := cntrl.Service.AddNewBook(c.Request.Context(), form)
	c.JSON(200, utils.WebResponses(c.Writer, 200, "success add new data", dataForm))

}

func (cntrl *LibraryControllerImpl) UpdateBook(c *gin.Context) {
	idBook := c.Param("id")
	form := web.UpdateBook{}
	idInt, _ := strconv.Atoi(idBook)
	form.Book_id = idInt

	if errBind := c.BindJSON(&form); errBind != nil {
		utils.LogErrorWithPanic(errBind)
	}

	successEdit, errEdit := cntrl.Service.UpdateBook(c.Request.Context(), form)

	if errEdit != nil {
		c.JSON(200, utils.WebResponses(c.Writer, 400, "failed update book", nil))
		return
	}

	c.JSON(200, utils.WebResponses(c.Writer, 200, "success update book", successEdit))

}

func (cntrl *LibraryControllerImpl) DeleteBook(c *gin.Context) {
	idBook := c.Param("id")
	idInt, _ := strconv.Atoi(idBook)

	dataBook, errNotFound := cntrl.Service.FindByIdBook(c.Request.Context(), idInt)

	if errNotFound != nil || dataBook.Book == "" {
		c.JSON(200, utils.WebResponses(c.Writer, 400, "failed delete book, book not found", nil))
		return
	}

	successDelete := cntrl.Service.DeleteBook(c.Request.Context(), idInt)
	c.JSON(200, utils.WebResponses(c.Writer, 200, "success delete data", successDelete))

}

func (cntrl *LibraryControllerImpl) FindBookById(c *gin.Context) {
	idBoook := c.Param("id")
	idInt, _ := strconv.Atoi(idBoook)
	response, err := cntrl.Service.FindByIdBook(c.Request.Context(), idInt)

	if err != nil {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "book not found", nil))
		return

	}

	c.JSON(200, utils.WebResponses(c.Writer, 200, "success get book", response))

}

func (cntrl *LibraryControllerImpl) FindAuthorById(c *gin.Context) {
	idAuthor := c.Param("id")
	idInt, _ := strconv.Atoi(idAuthor)
	response, err := cntrl.Service.FindByIdAuthor(c.Request.Context(), idInt)

	if err != nil || response.Author == "" {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "author not found", nil))
		return

	}

	c.JSON(200, utils.WebResponses(c.Writer, 200, "success get author", response))

}

func (cntrl *LibraryControllerImpl) FindAllBook(c *gin.Context) {
	booksData := cntrl.Service.FindAllBook(c.Request.Context())
	c.JSON(200, utils.WebResponses(c.Writer, 200, "success get all books", booksData))
}

func (cntrl *LibraryControllerImpl) FindAuthorByName(c *gin.Context) {
	authorName := c.Query("author")
	get, err := cntrl.Service.FindByNameAuthor(c.Request.Context(), authorName)

	if err != nil {
		c.JSON(404, utils.WebResponses(c.Writer, 404, "author not found", err))
		return
	}

	c.JSON(200, utils.WebResponses(c.Writer, 200, "success get author", get))

}

func (cntrl *LibraryControllerImpl) FindAllAuthorWithTheBook(c *gin.Context) {
	authorsWithBooks := cntrl.Service.GetAllAuthor(c.Request.Context())
	c.JSON(200, utils.WebResponses(c.Writer, 200, "success get all author", authorsWithBooks))
}
