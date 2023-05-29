package controller

import (
	"fmt"
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
	errBind := c.BindJSON(&form)
	if errBind != nil {
		utils.LogErrorWithPanic(errBind)
	}

	if form.Author == "" || form.Book_Title == "" {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "failed add data, ", nil))
	} else {
		dataForm := cntrl.Service.AddNewBook(c.Request.Context(), form)
		c.JSON(200, utils.WebResponses(c.Writer, 200, "success post data", dataForm))
	}

}

func (cntrl *LibraryControllerImpl) UpdateBook(c *gin.Context) {
	idBook := c.Param("id")
	form := web.UpdateBook{}
	idInt, _ := strconv.Atoi(idBook)
	form.Book_id = idInt
	c.BindJSON(&form)
	successEdit, errEdit := cntrl.Service.UpdateBook(c.Request.Context(), form)
	if errEdit != nil {
		fmt.Println(errEdit)
		c.JSON(200, utils.WebResponses(c.Writer, 400, "failed edit data, ", nil))
	} else {
		c.JSON(200, utils.WebResponses(c.Writer, 200, "success edit data", successEdit))
	}

}

func (cntrl *LibraryControllerImpl) DeleteBook(c *gin.Context) {
	idBook := c.Param("id")
	idInt, _ := strconv.Atoi(idBook)

	dataBook, errNotFound := cntrl.Service.FindByIdBook(c.Request.Context(), idInt)
	if errNotFound != nil || dataBook.Book == "" {
		c.JSON(200, utils.WebResponses(c.Writer, 400, "data not found", nil))
	} else {
		successDelete := cntrl.Service.DeleteBook(c.Request.Context(), idInt)
		c.JSON(200, utils.WebResponses(c.Writer, 200, "success delete data", successDelete))
	}
}

func (cntrl *LibraryControllerImpl) FindBookById(c *gin.Context) {
	idBoook := c.Param("id")
	idInt, _ := strconv.Atoi(idBoook)
	response, err := cntrl.Service.FindByIdBook(c.Request.Context(), idInt)
	if err != nil {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "data not found", nil))

	} else {
		c.JSON(200, utils.WebResponses(c.Writer, 200, "success get data", response))
	}
}

func (cntrl *LibraryControllerImpl) FindAuthorById(c *gin.Context) {
	idAuthor := c.Param("id")
	idInt, _ := strconv.Atoi(idAuthor)
	response, err := cntrl.Service.FindByIdAuthor(c.Request.Context(), idInt)
	if err != nil || response.Author == "" {
		c.JSON(400, utils.WebResponses(c.Writer, 400, "data not found", nil))

	} else {
		c.JSON(200, utils.WebResponses(c.Writer, 200, "success get data", response))
	}
}

func (cntrl *LibraryControllerImpl) FindAllBook(c *gin.Context) {
	booksData := cntrl.Service.FindAllBook(c.Request.Context())
	c.JSON(200, web.WebResponse{Status: 200, Message: "success", Data: booksData})
}

func (cntrl *LibraryControllerImpl) FindAuthorByName(c *gin.Context) {
	authorName := c.Query("author")
	get, err := cntrl.Service.FindByNameAuthor(c.Request.Context(), authorName)

	if err != nil {
		c.JSON(404, utils.WebResponses(c.Writer, 404, "author not found", err))
	} else {
		c.JSON(200, utils.WebResponses(c.Writer, 200, "author found", get))
	}
}
