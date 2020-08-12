package httputils

// Response basic type for API responses.
type Response struct {
    StatusCode int
    Body string
}

// NotFoundResponse default resource not found response.
var NotFoundResponse Response = Response{ 404, "Error: Resource not found."}