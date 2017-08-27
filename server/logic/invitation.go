package logic

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/logger"
	"ChineseChess/server/models"
	modelCache "ChineseChess/server/models/cache"
)

var (
	invitees map[string]chan string = make(map[string]chan string) // 受邀者
	invitors map[string]chan bool   = make(map[string]chan bool)   // 邀请者
)

// 消费邀请类型
type InvitationReplyOp = bool

const (
	InvitationReplyOpAccept InvitationReplyOp = true  // 接受
	InvitationReplyOpReject InvitationReplyOp = false // 拒绝
)

// InvitationForm contains invitor and invitee of an invitation
type InvitationForm struct {
	Invitor string `json:"invitor"` // 邀请者
	Invitee string `json:"invitee"` // 受邀者
}

// InvitationResp contains all info including 'op' of an invitaion
type InvitationResp struct {
	Invitor string            `json:"invitor"` // 邀请者
	Invitee string            `json:"invitee"` // 受邀者
	Op      InvitationReplyOp `json:"op"`      // 对邀请的回复
}

// TODO fix parameters and unmarshal the msg
// Invite invites a user to play game
func Invite(invitor, invitee string) (*models.ChessBoard, error) {

	// 判断受邀者状态
	session, err := modelCache.FindSession(invitee)
	if err != nil {
		return nil, errors.New("对方不在线")
	}
	if session.Status != modelCache.SessionStatusOK {
		return nil, errors.New("对方拒绝邀请")
	}

	if bson.IsObjectIdHex(invitor) && bson.IsObjectIdHex(invitee) {

		invitation := models.NewInvitation(bson.ObjectIdHex(invitor), bson.ObjectIdHex(invitee))
		if err := daf.Insert(invitation); err != nil {
			return nil, err
		}

		// 通知受邀者
		timeout := make(chan bool)
		go func() { invitees[invitee] <- invitor }()
		go func() { time.Sleep(30 * time.Second); timeout <- true }()
		select {
		case res := <-invitors[invitor]:

			close(timeout)

			// 对方回应(理论上邀请者每个等待回复期间只邀请了一个玩家)
			if res {

				// 对方接受邀请
				board := models.NewChessBoard(invitor, invitee)
				if err := daf.Insert(board); err != nil {
					return nil, errors.New("服务器出错")
				}
				if err := cache.AddBoardCache(board); err != nil {
					return nil, errors.New("服务器出错")
				}
				// (理论上此处不会并发竞争)
				if err := modelCache.UpdateSession(invitor, modelCache.SessionFieldStatus, modelCache.SessionStatusGame); err != nil {
					logger.Warn(err)
				}
				return board, nil
			}
			return nil, errors.New("对方拒绝邀请")
		case <-timeout:

			// 邀请超时
			close(timeout)
			return nil, errors.New("对方未接受")
		}

		return nil, nil
	}
	return nil, errors.New("数据不合法")
}

// TODO fix
// ReplyInvitation will reply an invitation, either accepting or rejection
func ReplyInvitation(invitor, invitee string, op InvitationReplyOp) (*models.ChessBoard, error) {

	if op == InvitationReplyOpReject {

		// 拒绝
		session, err := modelCache.FindSession(invitor)

		// 邀请者正在等待回复
		// TODO 并发安全处理
		if err == nil && session.Status == modelCache.SessionStatusOK {
			invitors[invitor] <- false
		}
		return nil, errors.New("已拒绝")
	}

	session, err := modelCache.FindSession(invitor)
	if err != nil || session.Status != modelCache.SessionStatusOK {
		return nil, errors.New("邀请已过期")
	}
	invitors[invitor] <- true
	return nil, nil
}

// JoinInvitees helps a user joining the invitees
func JoinInvitees(userID string) {
	if _, ok := invitees[userID]; ok {
		close(invitees[userID])
		delete(invitees, userID)
	}
	invitees[userID] = make(chan string)
}

// ExitInvitees helps a user to exit the invitees
func ExitInvitees(userID string) {
	if _, ok := invitees[userID]; ok {
		close(invitees[userID])
		delete(invitees, userID)
	}
}

// JoinInvitors helps a user joining the invitors
func JoinInvitors(userID string) {
	if _, ok := invitors[userID]; ok {
		close(invitors[userID])
		delete(invitors, userID)
	}
	invitors[userID] = make(chan bool)
}

// ExitInvitors helps a user to exit the invitors
func ExitInvitors(userID string) {
	if _, ok := invitors[userID]; ok {
		close(invitors[userID])
		delete(invitors, userID)
	}
}
