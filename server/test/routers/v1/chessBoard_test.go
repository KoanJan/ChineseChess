package v1

import (
	"encoding/json"
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "ChineseChess/server/routers/common"
	"ChineseChess/server/routers/v1"
)

type chessBoardForm struct {
	RedUserID   string `json:"red_user_id"`
	BlackUserID string `json:"black_user_id"`
}

func Test_CreateChessBoard(t *testing.T) {

	form := new(chessBoardForm)
	form.RedUserID = bson.NewObjectId().Hex()
	form.BlackUserID = bson.NewObjectId().Hex()
	data, _ := json.Marshal(form)
	t.Log(string(data))
	res_bytes := v1.CreateChessBoard(data)

	res := new(Resp)
	err := json.Unmarshal(res_bytes, res)
	if err != nil {
		t.Error(err)
	} else if res.Status != RespStatusOK {
		t.Error(res.Error)
	}
}
