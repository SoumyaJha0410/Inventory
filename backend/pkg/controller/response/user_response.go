package response

type ErrorResponse struct{
	ErrorResponse string `json:"error_message"`
}

func NewErrorResponse(errorMessage string) *ErrorResponse{
	return &ErrorResponse{errorMessage}
}
