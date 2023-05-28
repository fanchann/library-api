package utils

import (
	"fanchann/library/internal/models/domain"
	"fanchann/library/internal/models/web"
)

func AddNewBookToDomainAuthor(author *web.AddNewBooks) *domain.Author {
	return &domain.Author{Author_Name: author.Author}
}

func AddNewBookToDomainBook(book *web.AddNewBooks) *domain.Books {
	return &domain.Books{Book_Title: book.Book_Title}

}

func AuthorAndBookDomainToBookInfosDomain(author *domain.Author, book *domain.Books) *domain.Books_Information {
	return &domain.Books_Information{Book_id: book.Book_id, Author_id: author.Author_Id}
}

func UpdateBookToDomainBook(book *web.UpdateBook) *domain.Books {
	return &domain.Books{Book_Title: book.Book_Title}
}

func UpdateAuthorToDomainBook(book *web.UpdateBook) *domain.Author {
	return &domain.Author{Author_Name: book.Author}
}
