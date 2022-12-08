package results

import "github.com/shiqiyue/gof/resultcode"

// 返回结果
type Result struct {
	Msg  string      `json:"msg"`
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func Suc(data interface{}) *Result {
	return &Result{
		Msg:  SUCCESS_MSG,
		Code: resultcode.SUCCESS,
		Data: data,
	}
}

func Err(err error) *Result {
	return &Result{
		Msg:  err.Error(),
		Code: resultcode.FAIL,
		Data: nil,
	}

}
