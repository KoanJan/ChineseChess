package render

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// 返回的数据格式
type respData struct {
	Code int                    `json:"code"` // http状态码
	Msg  string                 `json:"msg"`  // 提示信息
	Data map[string]interface{} `json:"data"` // 数据
}

// RenderErr renders a error
func RenderErr(c *gin.Context, err error, code ...int) {

	msg := ""
	if err != nil {
		msg = err.Error()
	}
	_code := 500
	if len(code) > 0 {
		_code = code[0]
	}
	render(c, _code, msg, nil)
}

// RenderOk renders a ok status and some data if existing
func RenderOk(c *gin.Context, data ...map[string]interface{}) {

	if len(data) > 0 {
		render(c, 200, "ok", data[0])
	} else {
		render(c, 200, "ok", nil)
	}
}

func render(c *gin.Context, code int, msg string, data map[string]interface{}) {
	if data == nil {
		data = make(map[string]interface{})
	}
	b, _ := json.Marshal(respData{code, msg, data})
	c.Writer.Write(b)
	c.AbortWithStatus(200)
}
