package logic

import (
	"errors"
	"fmt"

	"ChineseChess/server/routers/ws/msg"
)

const (
	GameLogicFuncPlay  = "play"  // 下子
	GameLogicFuncMatch = "match" // 匹配
)

var msgBox chan *msg.GameServerMsg = make(chan *msg.GameServerMsg, 1024) // 消息推送

// 推送消息
func SendGameServerMsg(call string, data []byte, err error, uid ...string) {

	gameServerMsg := new(msg.GameServerMsg)
	gameServerMsg.UIDs = uid // broadcast if len(uid) is 0
	gameMsg := new(msg.GameMsg)
	gameMsg.Call = call
	gameMsg.Data = data
	if err != nil {
		gameMsg.Err = err.Error()
	}
	gameServerMsg.GameMsg = gameMsg

	msgBox <- gameServerMsg
}

// 监听服务端产生的游戏消息
func HandleGameServerMsg(handler func(*msg.GameServerMsg)) {

	for {
		select {
		case gsm := <-msgBox:
			go handler(gsm)
		}
	}
}

var api map[string]func(*msg.GameMsg, ...string) = map[string]func(*msg.GameMsg, ...string){

	GameLogicFuncPlay: Play,

	GameLogicFuncMatch: Match,
}

// GameLogicFunc returns a logic func if exists,
// or else a func that always returns a error says that the function doesn't exist.
func GameLogicFunc(funcName string) func(gameMsg *msg.GameMsg, uid ...string) {

	if f, existed := api[funcName]; existed {
		return f
	}
	return noSuchFunc(funcName)
}

func noSuchFunc(funcName string) func(*msg.GameMsg, ...string) {

	return func(gameMsg *msg.GameMsg, uid ...string) {

		SendGameServerMsg(gameMsg.Call, []byte{}, errors.New(fmt.Sprintf("no such func called '%s'", string(funcName))), uid[0])
	}
}
