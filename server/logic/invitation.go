package logic

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/logger"
	"ChineseChess/server/models"
	modelCache "ChineseChess/server/models/cache"
	"ChineseChess/server/routers/ws/msg"
)

var (
	invitationMsgs map[string]*InvitationMsg          = make(map[string]*InvitationMsg)          // 维护所有邀请函状态
	inviteeBoxes   map[string]chan *InvitationMsg     = make(map[string]chan *InvitationMsg)     // 受邀者
	invitorBoxes   map[string]chan *InvitationMsgResp = make(map[string]chan *InvitationMsgResp) // 邀请者
)

// 邀请函简要
type InvitationMsg struct {
	InvitationID string `json:"invitation_id"` // 邀请函id
	Invitor      string `json:"invitor"`       // 邀请者
	Invitee      string `json:"invitee"`       // 受邀者
}

// 转化成Model
func (this *InvitationMsg) AsModel() *models.Invitation {
	invitation := new(models.Invitation)
	invitation.ID = bson.ObjectIdHex(this.InvitationID)
	invitation.Invitor = bson.ObjectIdHex(this.Invitor)
	invitation.Invitee = bson.ObjectIdHex(this.Invitee)
	return invitation
}

// 邀请函回复
type InvitationMsgResp struct {
	InvitationID string `json:"invitation_id"` // 邀请函id
	Accept       bool   `json:"accept"`        // 接受
}

// 邀请表单
type InvitationForm struct {
	Invitor string `json:"invitor"` // 邀请者
	Invitee string `json:"invitee"` // 受邀者
}

// 撤销邀请的表单
type InvitationCancelForm struct {
	InvitationID string `json:"invitation_id"` // 邀请函id
}

// 邀请结果
type InvitationResult struct {
	Board *models.ChessBoard `json:"board"`
	Error string             `json:"error"`
}

// 发出邀请
func Invite(gameMsg *msg.GameMsg, invitor string) {

	form := new(InvitationForm)
	data, err := proto.Marshal(gameMsg)
	if err != nil {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("数据解析失败"), invitor)
		return
	}
	if err = json.Unmarshal(data, form); err != nil {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("数据解析失败"), invitor)
		return
	}

	// 判断受邀者状态
	session, err := modelCache.FindSession(form.Invitee)
	if err != nil {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("对方不在线"), invitor)
		return
	}
	if session.Status != modelCache.SessionStatusOK {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("对方拒绝邀请"), invitor)
		return
	}

	if bson.IsObjectIdHex(invitor) && bson.IsObjectIdHex(form.Invitee) {

		// 暂存邀请函
		invitationMsg := &InvitationMsg{bson.NewObjectId().Hex(), invitor, form.Invitee}
		invitationMsgs[invitationMsg.InvitationID] = invitationMsg

		// 通知受邀者
		timeout := make(chan bool)
		go func() {
			if box, ok := inviteeBoxes[form.Invitee]; ok {
				box <- invitationMsg
			}
		}()
		go func() { time.Sleep(30 * time.Second); timeout <- true }()
		select {
		case msg := <-invitorBoxes[invitor]:

			close(timeout)

			// 对方回应(理论上邀请者每个等待回复期间只邀请了一个玩家)
			if msg.Accept {

				// 自己已经取消邀请
				if _, ok := invitationMsgs[invitationMsg.InvitationID]; !ok {
					return
				}

				// 对方接受邀请
				board := models.NewChessBoard(invitor, form.Invitee)
				if err := daf.Insert(board); err != nil {
					PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("服务器出错"), invitor, form.Invitee)
					return
				}
				if err := cache.AddBoardCache(board); err != nil {
					PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("服务器出错"), invitor, form.Invitee)
					return
				}
				invitaion := invitationMsg.AsModel()
				invitaion.ChessBoardID = board.ID.Hex()
				if err := daf.Insert(invitaion); err != nil {
					PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("服务器出错"), invitor, form.Invitee)
					return
				}

				// (理论上此处不会并发竞争)
				if err := modelCache.UpdateSession(invitor, modelCache.SessionFieldStatus, modelCache.SessionStatusGame); err != nil {
					logger.Warn(err)
				}
				data, _ := json.Marshal(board)
				PushGameServerMsg(GameLogicFuncInvite, data, nil, invitor, form.Invitee)
				return
			}
			PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("对方拒绝邀请"), invitor)
			return
		case <-timeout:

			// 邀请超时
			close(timeout)
			delete(invitationMsgs, invitationMsg.InvitationID)
			PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("对方未接受"), invitor)
			return
		}
	}
	PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("数据不合法"), invitor)
	return
}

// 回复邀请
func ReplyInvitation(gameMsg *msg.GameMsg, invitee string) {

	resp := new(InvitationMsgResp)
	data, err := proto.Marshal(gameMsg)
	if err != nil {
		PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("数据解析失败"), invitee)
		return
	}
	if err = json.Unmarshal(data, resp); err != nil {
		PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("数据解析失败"), invitee)
		return
	}

	invitaion, ok := invitationMsgs[resp.InvitationID]
	if !ok {
		PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("邀请已取消"), invitee)
		return
	}

	// 检查邀请者是否在线
	session, err := modelCache.FindSession(invitaion.Invitor)

	if err != nil {

		// 邀请者不在线
		PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("对方不在线"), invitee)
		return
	}

	if session.Status != modelCache.SessionStatusOK {

		// 邀请者正忙
		PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("对方忙"), invitee)
		return
	}

	if box, ok := invitorBoxes[invitaion.InvitationID]; ok {

		box <- resp
		return
	}

	PushGameServerMsg(GameLogicFuncReplyInvitation, []byte{}, errors.New("对方不在线"), invitee)
}

// 取消邀请
func CancelInvitation(gameMsg *msg.GameMsg, invitor string) {

	form := new(InvitationCancelForm)
	data, err := proto.Marshal(gameMsg)
	if err != nil {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("数据解析失败"), invitor)
		return
	}
	if err = json.Unmarshal(data, form); err != nil {
		PushGameServerMsg(GameLogicFuncInvite, []byte{}, errors.New("数据解析失败"), invitor)
		return
	}

	delete(invitationMsgs, form.InvitationID)
}

// AddUserIntoInvitationSystem adds a new user into invitation system
func AddUserIntoInvitationSystem(userID string) {

	DeleteUserFromInvitationSystem(userID)

	// invitee
	inviteeBoxes[userID] = make(chan *InvitationMsg)

	// invitor
	invitorBoxes[userID] = make(chan *InvitationMsgResp)
}

// DeleteUserFromInvitationSystem delete a user from invitation system
func DeleteUserFromInvitationSystem(userID string) {

	// invitee
	if _, ok := inviteeBoxes[userID]; ok {
		close(inviteeBoxes[userID])
		delete(inviteeBoxes, userID)
	}

	// invitor
	if _, ok := invitorBoxes[userID]; ok {
		close(invitorBoxes[userID])
		delete(invitorBoxes, userID)
	}
}

// 初始化
func initInvitation() {

}
