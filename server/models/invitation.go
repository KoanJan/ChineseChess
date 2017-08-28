package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*
邀请
*/
type Invitation struct {
	Common
	Invitor      bson.ObjectId `bson:"invitor",json:"invitor"`               // 邀请者
	Invitee      bson.ObjectId `bson:"invitee",json:"invitee"`               // 受邀者
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
	invitation.CreatedAt = time.Now()
	invitation.UpdatedAt = time.Now()
	return invitation
}
