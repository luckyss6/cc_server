package handler

const (
	SuccessCode = iota + 1000
)

const (
	ErrCode = iota + 2000
	BindingParamErrCode
	LoginErrCode
	DBErrCode
	TokenErrCode
	DockerErrCode
)

const (
	BindingParamErrMsg = "参数绑定错误"
	DBErrMsg           = "数据库错误"
	LoginErrMsg        = "用户名或密码错误"
	TokenErrMsg        = "token错误"
	DockerErrMsg       = "docker错误"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
func success(data interface{}) *Response {
	return NewResponse(SuccessCode, "success", data)
}

func fail(code int, msg string) *Response {
	return NewResponse(code, msg, nil)
}
