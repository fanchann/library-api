package services

import (
	"context"
	"database/sql"
	"fmt"

	"fanchann/library/internal/models/domain"
	"fanchann/library/internal/models/web"
	"fanchann/library/internal/repositories/authors"
	"fanchann/library/internal/repositories/books"
	booksinformation "fanchann/library/internal/repositories/books_information"
	"fanchann/library/pkg/utils"
)

var (
	nullAuthor = domain.Author{}
)

type LibraryImpl struct {
	DB           *sql.DB
	AuthorRepo   authors.IAuthorsRepositories
	BookRepo     books.IBooksRepositories
	BookInfoRepo booksinformation.IBooksInformation
}

func NewLibraryImpl(db *sql.DB, authorRepo authors.IAuthorsRepositories, bookRepo books.IBooksRepositories, bookInfos booksinformation.IBooksInformation) ILibraryServices {
	return &LibraryImpl{DB: db, AuthorRepo: authorRepo, BookRepo: bookRepo, BookInfoRepo: bookInfos}
}

func (lib *LibraryImpl) AddNewBook(ctx context.Context, formData web.AddNewBooks) web.AddNewBooks {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	// find author first
	authorFound := lib.AuthorRepo.FindAuthorByName(ctx, tx, formData.Author)

	// check if author found
	if authorFound.Author_Name != "" {
		bookData := lib.BookRepo.Add(ctx, tx, utils.AddNewBookToDomainBook(&formData))
		lib.BookInfoRepo.Insert(ctx, tx, utils.AuthorAndBookDomainToBookInfosDomain(&domain.Author{Author_Id: authorFound.Author_Id}, &bookData))

		return web.AddNewBooks{Author: formData.Author, Book_Title: bookData.Book_Title}
	}

	authorData := lib.AuthorRepo.Add(ctx, tx, utils.AddNewBookToDomainAuthor(&formData))
	bookData := lib.BookRepo.Add(ctx, tx, utils.AddNewBookToDomainBook(&formData))
	lib.BookInfoRepo.Insert(ctx, tx, utils.AuthorAndBookDomainToBookInfosDomain(&authorData, &bookData))

	return web.AddNewBooks{Author: authorData.Author_Name, Book_Title: bookData.Book_Title}
}

func (lib *LibraryImpl) UpdateBook(ctx context.Context, formData web.UpdateBook) (web.UpdateBook, error) {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	// before update check data book is exist?
	dataBook, errBookNotFound := lib.BookRepo.FindById(ctx, tx, formData.Book_id)

	if errBookNotFound != nil {
		return web.UpdateBook{}, errBookNotFound

	}

	// don't forget to check author name
	authorFound := lib.AuthorRepo.FindAuthorByName(ctx, tx, formData.Author)

	// if author exist?
	if authorFound == nullAuthor {
		authorData := lib.AuthorRepo.Add(ctx, tx, utils.UpdateAuthorToDomainBook(&formData))
		lib.BookInfoRepo.Delete(ctx, tx, dataBook.Book_id)
		lib.BookInfoRepo.Insert(ctx, tx, &domain.Books_Information{Book_id: formData.Book_id, Author_id: authorData.Author_Id})
		domainBook := lib.BookRepo.Update(ctx, tx, utils.UpdateBookToDomainBook(&formData))

		return web.UpdateBook{Book_id: domainBook.Book_id, Book_Title: domainBook.Book_Title, Author: formData.Author}, nil
	}

	return web.UpdateBook{}, utils.ErrorWithReturn("failed edit data")
}

func (lib *LibraryImpl) DeleteBook(ctx context.Context, bookID int) error {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	// before delete check data book first
	data, errBookNotFound := lib.BookRepo.FindById(ctx, tx, bookID)

	if errBookNotFound != nil {
		return utils.ErrorWithReturn(fmt.Sprintf("book with id %d not found", bookID))
	}

	// if exist, delete the book
	errBookRepo := lib.BookInfoRepo.Delete(ctx, tx, data.Book_id)
	lib.BookRepo.Delete(ctx, tx, data.Book_id)
	utils.LogErrorWithPanic(errBookRepo)

	return nil
}

func (lib *LibraryImpl) FindByIdBook(ctx context.Context, bookID int) (web.BooksResponse, error) {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	dataBookWithAuthor, errNotFound := lib.BookInfoRepo.FindBookIdWithAuthor(ctx, tx, bookID)

	if errNotFound != nil || dataBookWithAuthor.Author_Name == "" {
		return web.BooksResponse{}, errNotFound
	}

	return web.BooksResponse{Book: dataBookWithAuthor.Book_Title, Author: dataBookWithAuthor.Author_Name}, nil
}

func (lib *LibraryImpl) FindByIdAuthor(ctx context.Context, authorID int) (web.AuthorsResponse, error) {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	authorFound, errAuthorNotFound := lib.AuthorRepo.FindAuthorById(ctx, tx, authorID)

	if errAuthorNotFound != nil || authorFound.Author_Name == "" {
		return web.AuthorsResponse{}, errAuthorNotFound
	}

	dataBooks := lib.BookInfoRepo.FindAuthorIdWithBooks(ctx, tx, authorID)
	books := []string{}

	for _, book := range dataBooks {
		books = append(books, book.Book_Title)
	}

	return web.AuthorsResponse{Author: authorFound.Author_Name, Books: books}, nil

}

func (lib *LibraryImpl) FindAllBook(ctx context.Context) []web.BooksResponse {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	var booksResponses []web.BooksResponse

	dataBooksAndAuthors := lib.BookInfoRepo.FindAllBooksAndAuthor(ctx, tx)

	for _, authorBook := range dataBooksAndAuthors {
		bookAndAuthor := utils.BookAndAuthorDomainToWeb(&authorBook)
		booksResponses = append(booksResponses, bookAndAuthor)
	}

	return booksResponses
}

func (lib *LibraryImpl) FindByNameAuthor(ctx context.Context, authorName string) (web.AuthorsResponse, error) {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	authorData := lib.AuthorRepo.FindAuthorByName(ctx, tx, authorName)
	if authorData == nullAuthor {
		return web.AuthorsResponse{}, utils.ErrorWithReturn(fmt.Sprintf("author with name %s not found", authorName))
	}

	booksWithAuthor := lib.BookInfoRepo.FindAuthorIdWithBooks(ctx, tx, authorData.Author_Id)
	books := []string{}

	for _, book := range booksWithAuthor {
		books = append(books, book.Book_Title)
	}

	return web.AuthorsResponse{Author: authorData.Author_Name, Books: books}, nil
}

func (lib *LibraryImpl) GetAllAuthor(ctx context.Context) []web.AuthorsResponse {
	tx := utils.StartTransaction(lib.DB)
	defer utils.TransactionsCommitOrRollback(tx)

	authors := lib.AuthorRepo.FindAllAuthor(ctx, tx)

	var allAuthors []web.AuthorsResponse
	for _, authorValue := range authors {
		var author web.AuthorsResponse
		authorBooks := lib.BookInfoRepo.FindAuthorIdWithBooks(ctx, tx, authorValue.Author_Id)
		author.Author = authorValue.Author_Name
		author.Books = utils.DomainBookListToSlice(authorBooks)

		allAuthors = append(allAuthors, author)
	}

	return allAuthors
}
