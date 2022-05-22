package response

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty" `
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
}

func FailResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

