package helper

type ApiError struct {
	Error   error
	Message string
	Code    int
}

func ErrorHandler (apiError ApiError) *ApiError {
	if apiError.Code != 0 {
		return &apiError
	}
	apiError.Code = 500
	return &apiError
}