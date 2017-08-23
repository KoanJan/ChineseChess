package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 邀请状态
const (
	_                        = iota
	InvitationStatusInvited  // 邀请
	InvitationStatusAccepted // 接受
	InvitationStatusRejected // 拒绝
)

/*
邀请
*/
type Invitation struct {
	Common
	Invitor      bson.ObjectId `bson:"invitor",json:"invitor"`               // 邀请者
	Invitee      bson.ObjectId `bson:"invitee",json:"invitee"`               // 受邀者
	Status       int32         `bson:"status",json:"status"`                 // 邀约状态
	ChessBoardID string        `bson:"chess_board_id",json:"chess_board_id"` // 棋局ID
}

func (this *Invitation) CollectionName() string {
	return "invitation"
}

/*
发出新邀请
*/
func NewInvitation(invitor, invitee bson.ObjectId) *Invitation {

	invitation := new(Invitation)
	invitation.ID = bson.NewObjectId()
	invitation.Invitor = invitor
	invitation.Invitee = invitee
	invitation.Status = InvitationStatusInvited
	invitation.CreatedAt = time.Now()
	invitation.UpdatedAt = time.Now()
	return invitation
}
