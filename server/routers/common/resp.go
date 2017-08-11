package common

import "encoding/json"

// 请求结果状态
const (
	RespStatusOK    = iota // 成功
	RespStatusError        // 失败
)

// 请求返回数据结构
type Resp struct {
	Status int         `json: "status"` // 状态
	Data   interface{} `json: "data"`   // 数据
	Msg    string      `json: "msg"`    // 提示信息
	Error  error       `json: "error"`  // 错误信息
}

// 成功
func RespOK() []byte {
	return resp(nil, "ok", nil)
}

// 成功(有数据)
func RespData(data interface{}) []byte {
	return resp(data, "ok", nil)
}

// 成功(自定义提示)
func RespOKWithMessage(data interface{}, msg string) []byte {
	return resp(data, msg, nil)
}

// 错误
func RespErr(err error) []byte {
	return resp(nil, "", err)
}

func resp(data interface{}, msg string, err error) []byte {

	var status int
	if err != nil {
		status = RespStatusError
	} else {
		status = RespStatusOK
	}
	resp := &Resp{status, data, msg, err}
	b, _ := json.Marshal(resp)
	return b
}
