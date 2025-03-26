package errorx

// DefaultCode 默认错误码
const DefaultCode = 1

// CodeError 定义了一个错误的结构体，包含错误代码和错误消息
type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CodeErrorResponse 定义了一个错误响应的结构体，用于对外部错误的统一响应格式
type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewCodeError 创建并返回一个新的CodeError实例
// 参数code: 错误代码
// 参数msg: 错误消息
func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

// NewDefaultError 创建并返回一个带有默认错误代码的CodeError实例
// 参数msg: 错误消息
func NewDefaultError(msg string) error {
	return NewCodeError(DefaultCode, msg)
}

// Error 实现了error接口，返回错误消息
func (e *CodeError) Error() string {
	return e.Message
}

// Data 返回CodeError的响应格式实例，便于统一错误响应格式
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}
