package httputils

// ErrorResponse basic type for API Error responses.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// APIResponse basic type for API responses.
type APIResponse struct {
	Data interface{} `json:"data"`
}

// NotFoundResponse default resource not found response.
var NotFoundResponse ErrorResponse = ErrorResponse{404, "Error: Resource not found."}

// NewAPIResponse default API response.
func NewAPIResponse(data interface{}) interface{} {
	// response := make(map[string]string)

	return APIResponse{
		Data: data,
	}
}
