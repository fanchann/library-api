package web

type (
	AddNewBooks struct {
		Book_name   string `json:"book_name"`
		Book_Author string `json:"book_author"`
	}

	UpdateBook struct {
		Book_id     int    `json:"book_id"`
		Book_name   string `json:"book_name"`
		Book_Author string `json:"book_author"`
	}
)
