package controller

var (
	User = &UserController{}
	Docker = &DockerController{}
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = iota + 1000
	InvalidParamCode
	GetDockerInfoFailCode
	CreateContainerFailCode
	DBErrorCode
	LoginFailCode
)

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(data interface{}) *Response {
	return NewResponse(SuccessCode, "success", data)
}

func Fail(code int, msg string) *Response {
	return NewResponse(code, msg, nil)
}