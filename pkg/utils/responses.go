package utils

import "fanchann/library/internal/models/web"

func WebResponses(status int, message string, data interface{}) web.WebResponse {
	return web.WebResponse{Status: status, Message: message, Data: data}
}
