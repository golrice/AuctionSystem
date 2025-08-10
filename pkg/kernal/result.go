package kernal

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 默认出现错误时候的结果
type ErrorResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewErrorResult(code int, msg string) *ErrorResult {
	return &ErrorResult{
		Code: code,
		Msg:  msg,
	}
}

// 默认出现成功时候的结果
type SuccessResult Result

func NewSuccessResult(data interface{}) *SuccessResult {
	return &SuccessResult{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func NewDefaultSuccessResult() *SuccessResult {
	return &SuccessResult{
		Code: 0,
		Msg:  "success",
		Data: struct{}{},
	}
}
