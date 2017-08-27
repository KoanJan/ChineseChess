package logic

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"ChineseChess/server/cache"
	"ChineseChess/server/daf"
	"ChineseChess/server/models"
)

var (
	invitees map[string]chan string = make(map[string]chan string) // 受邀者
	invitors map[string]chan bool   = make(map[string]chan bool)   // 邀请者
)

// InviteForm contains invitor and invitee of a invitation
type InviteForm struct {
	Invitor string `json:"invitor"` // 邀请者
	Invitee string `json:"invitee"` // 受邀者
}

// Invite invites a user to play game
func Invite(invitor, invitee string) (*models.ChessBoard, error) {

	// TODO 判断对方是否在游戏中, 如果游戏中则直接拒绝
	if !UserIsOnline(invitee) {
		return nil, errors.New("对方不在线")
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

			// 对方回应
			if res {
				// 对方接受邀请
				board := models.NewChessBoard(invitor, invitee)
				if err := daf.Insert(board); err != nil {
					return nil, errors.New("服务器出错")
				}
				if err := cache.AddBoardCache(board); err != nil {
					return nil, errors.New("服务器出错")
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
