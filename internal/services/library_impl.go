package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"fanchann/library/internal/models/domain"
	"fanchann/library/internal/models/web"
	"fanchann/library/internal/repositories/authors"
	"fanchann/library/internal/repositories/books"
	booksinformation "fanchann/library/internal/repositories/books_information"
	"fanchann/library/pkg/utils"
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
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	authorData := lib.AuthorRepo.Add(ctx, tx, utils.AddNewBookToDomainAuthor(&formData))
	bookData := lib.BookRepo.Add(ctx, tx, utils.AddNewBookToDomainBook(&formData))
	lib.BookInfoRepo.Insert(ctx, tx, utils.AuthorAndBookDomainToBookInfosDomain(&authorData, &bookData))

	return web.AddNewBooks{Author: authorData.Author_Name, Book_Title: bookData.Book_Title}
}

func (lib *LibraryImpl) UpdateBook(ctx context.Context, formData web.UpdateBook) (web.UpdateBook, error) {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)
	fmt.Println(formData)

	dataBook, errBookNotFound := lib.BookRepo.FindById(ctx, tx, formData.Book_id)
	if errBookNotFound != nil {
		return web.UpdateBook{}, errBookNotFound

	}

	authorFound := lib.AuthorRepo.FindAuthorByName(ctx, tx, formData.Author)

	fmt.Println("author -> ", authorFound)
	nullAuthor := domain.Author{}
	if authorFound == nullAuthor {
		authorData := lib.AuthorRepo.Add(ctx, tx, utils.UpdateAuthorToDomainBook(&formData))
		lib.BookInfoRepo.Delete(ctx, tx, dataBook.Book_id)
		lib.BookInfoRepo.Insert(ctx, tx, &domain.Books_Information{Book_id: formData.Book_id, Author_id: authorData.Author_Id})
		domainBook := lib.BookRepo.Update(ctx, tx, utils.UpdateBookToDomainBook(&formData))

		return web.UpdateBook{Book_id: domainBook.Book_id, Book_Title: domainBook.Book_Title, Author: formData.Author}, nil
	}
	return web.UpdateBook{}, errors.New("failed edit data")
}

func (lib *LibraryImpl) DeleteBook(ctx context.Context, bookID int) error {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	data, errBookNotFound := lib.BookRepo.FindById(ctx, tx, bookID)

	if errBookNotFound != nil {
		return err
	}

	errBookRepo := lib.BookInfoRepo.Delete(ctx, tx, data.Book_id)
	lib.BookRepo.Delete(ctx, tx, data.Book_id)
	utils.LogErrorWithPanic(errBookRepo)

	return nil
}

func (lib *LibraryImpl) FindByIdBook(ctx context.Context, bookID int) (web.BooksResponse, error) {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	dataBookWithAuthor, errNotFound := lib.BookInfoRepo.FindBookIdWithAuthor(ctx, tx, bookID)
	if err != nil || dataBookWithAuthor.Author_Name == "" {
		return web.BooksResponse{}, errNotFound
	}

	return web.BooksResponse{Book: dataBookWithAuthor.Book_Title, Author: dataBookWithAuthor.Author_Name}, nil
}

func (lib *LibraryImpl) FindByIdAuthor(ctx context.Context, authorID int) (web.AuthorsResponse, error) {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	authorFound, errAuthorNotFound := lib.AuthorRepo.FindAuthorById(ctx, tx, authorID)
	if errAuthorNotFound != nil || authorFound.Author_Name == "" {
		return web.AuthorsResponse{}, errAuthorNotFound
	}
	dataBooks := lib.BookInfoRepo.FindAuthorIdWithBooks(ctx, tx, authorID)
	books := []string{}
	for _, i := range dataBooks {
		books = append(books, i.Book_Title)
	}
	return web.AuthorsResponse{Author: authorFound.Author_Name, Books: books}, nil

}

func (lib *LibraryImpl) FindAllBook(ctx context.Context) []web.BooksResponse {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	var booksResponses []web.BooksResponse

	dataBooksAndAuthors := lib.BookInfoRepo.FindAllBooksAndAuthor(ctx, tx)
	for _, X := range dataBooksAndAuthors {
		bookAndAuthor := utils.BookAndAuthorDomainToWeb(&X)
		booksResponses = append(booksResponses, bookAndAuthor)
	}
	return booksResponses
}

func (lib *LibraryImpl) FindByNameAuthor(ctx context.Context, authorName string) (web.AuthorsResponse, error) {
	tx, err := lib.DB.Begin()
	utils.LogErrorWithPanic(err)
	defer utils.TransactionsCommitOrRollback(tx)

	nullAuthor := domain.Author{}

	authorData := lib.AuthorRepo.FindAuthorByName(ctx, tx, authorName)
	if authorData == nullAuthor {
		return web.AuthorsResponse{}, errors.New("gaono su")
	}
	booksWithAuthor := lib.BookInfoRepo.FindAuthorIdWithBooks(ctx, tx, authorData.Author_Id)
	books := []string{}
	for _, i := range booksWithAuthor {
		books = append(books, i.Book_Title)
	}
	return web.AuthorsResponse{Author: authorData.Author_Name, Books: books}, nil
}
