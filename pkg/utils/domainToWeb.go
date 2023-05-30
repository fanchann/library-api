package utils

import (
	"fanchann/library/internal/models/domain"
	"fanchann/library/internal/models/web"
)

func BookAndAuthorDomainToWeb(domain *domain.BookWithAuthor) web.BooksResponse {
	return web.BooksResponse{Author: domain.Author_Name, Book: domain.Book_Title}
}

func DomainBookListToSlice(book []domain.BookLists) []string {
	var books []string
	for _, buk := range book {
		books = append(books, buk.Book_Title)
	}
	return books
}
