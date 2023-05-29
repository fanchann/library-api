package utils

import (
	"net/http"

	"fanchann/library/internal/models/web"
)

func WebResponses(writer http.ResponseWriter,status int, message string, data interface{}) web.WebResponse {
	writer.WriteHeader(status)
	return web.WebResponse{Status: status, Message: message, Data: data}
}
