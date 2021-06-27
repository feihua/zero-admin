package shared

import (
	"google.golang.org/grpc/status"
	"net/http"
)

const defaultCode = 1001

func ErrorHandler(err error) (int, interface{}) {
	message := getGrpcErrorMessage(err)
	return http.StatusOK, CodeError{
		Code: -1,
		Msg:  message,
	}
}

func getGrpcErrorMessage(err error) string {
	s, flag := status.FromError(err)
	message := ""
	if flag {
		message = s.Message()
	} else {
		message = err.Error()
	}
	return message
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
