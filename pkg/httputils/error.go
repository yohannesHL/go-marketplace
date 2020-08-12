package httputils

import (
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// ServerError returns and handles default server error response.
func ServerError(err error) Response {
	logger.Println(err.Error())

	return Response{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}
}

// ClientError returns common client error response matching given statusCode.
func ClientError(statusCode int) Response {
	return Response{
		StatusCode: statusCode,
		Body:       http.StatusText(statusCode),
	}
}
