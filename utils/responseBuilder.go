package utils

import "github.com/iacopoghilardi/golang-backend-boilerplate/internals/types"

var (
	SuccessStatus = "OK"
	ErrorStatus   = "KO"
)

func BuildSuccessResponse(data interface{}) types.GenericSuccessResponse {
	return types.GenericSuccessResponse{
		Status: SuccessStatus,
		Data:   data,
	}
}

func BuildErrorResponse(message string, description string) types.GenericErrorResponse {
	return types.GenericErrorResponse{
		Status: ErrorStatus,
		Error: types.ErrorMessageResponse{
			Error:       message,
			Description: description,
		},
	}
}
