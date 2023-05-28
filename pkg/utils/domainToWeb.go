package utils

import (
	"fanchann/library/internal/models/domain"
	"fanchann/library/internal/models/web"
)

func BookAndAuthorDomainToWeb(domain *domain.BookWithAuthor) web.BooksResponse {
	return web.BooksResponse{Author: domain.Author_Name, Book: domain.Book_Title}
}
