package msg

// GameServerMsg is the msg type inner server
type GameServerMsg struct {
	UIDs    []string // UID
	GameMsg *GameMsg // 数据
}

// NewGameServerMsg
func NewGameServerMsg(gameMsg *GameMsg, uid ...string) *GameServerMsg {

	gameServerMsg := new(GameServerMsg)
	gameServerMsg.UIDs = uid
	gameServerMsg.GameMsg = gameMsg
	return gameServerMsg
}
