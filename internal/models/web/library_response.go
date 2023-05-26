package web

type (
	WebResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	AuthorsResponse struct {
		Author string      `json:"author"`
		Books  interface{} `json:"books"`
	}

	BooksResponse struct {
		Book   string `json:"book"`
		Author string `json:"author"`
	}
)
