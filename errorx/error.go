package errorx

import "encoding/json"

const ErrorCode = 500
const SuccessCode = 200
const SuccessMessage = "success"

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CodeErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CodeErrorResponseWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type CodeErrorWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewCodeErrorWithData(code int, msg string, data interface{}) error {
	error := &CodeErrorWithData{Code: code, Message: msg, Data: data}
	return error
}

func NewDefaultError(msg string) error {
	return NewCodeError(ErrorCode, msg)
}

func NewSuccessTip() error {
	return NewCodeError(SuccessCode, SuccessMessage)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeErrorWithData) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}
func (e *CodeError) JSON() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}
func (e *CodeErrorWithData) DataV2() *CodeErrorResponseWithData {
	return &CodeErrorResponseWithData{
		Code:    e.Code,
		Message: e.Message,
		Data:    e.Data,
	}
}
