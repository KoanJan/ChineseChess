package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 邀请
type Invitation struct {
	ID           bson.ObjectId `bson:"_id" json:"id,string"`                 // ID
	Invitor      bson.ObjectId `bson:"invitor" json:"invitor"`               // 邀请者
	Invitee      bson.ObjectId `bson:"invitee" json:"invitee"`               // 受邀者
	ChessBoardID string        `bson:"chess_board_id" json:"chess_board_id"` // 棋局ID
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`         // 创建时间
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`         // 修改时间
}

func (this *Invitation) CN() string {
	return InvitationCN()
}

func (this *Invitation) GetID() bson.ObjectId {
	return this.ID
}

func InvitationCN() string {
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
