package httputils

import (
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// ServerError returns and handles default server error response.
func ServerError(err error) ErrorResponse {
	logger.Println(err.Error())

	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
}

// ClientError returns common client error response matching given statusCode.
func ClientError(statusCode int) ErrorResponse {
	return ErrorResponse{
		Code:    statusCode,
		Message: http.StatusText(statusCode),
	}
}
