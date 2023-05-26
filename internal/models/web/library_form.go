package web

type (
	AddNewBooks struct {
		Book_Title string `json:"book_title"`
		Author     string `json:"author"`
	}

	UpdateBook struct {
		Book_id    int    `json:"book_id"`
		Book_Title string `json:"book_title"`
		Author     string `json:"author"`
	}
)
