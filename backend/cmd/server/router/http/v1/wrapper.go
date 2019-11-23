package v1

import (
	"SmartLocker/e"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func Wrap(err int, body interface{}) interface{} {
	if err != e.Success {
		return Response{
			Code: err,
			Msg:  e.GetMsg(err),
		}
	}
	return Response{
		Code: e.Success,
		Msg:  e.GetMsg(e.Success),
		Body: body,
	}
}
