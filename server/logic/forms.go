package logic

// PlayForm contains the all parameters that the Play API needs
type PlayForm struct {
	X1      int32  `json:"x1"`
	Y1      int32  `json:"y1"`
	X2      int32  `json:"x2"`
	Y2      int32  `json:"y2"`
	Step    int    `json:"step"` // 步数验证
	BoardID string `json:"board_id"`
	UserID  string `json:"user_id"`
}

// PlayResp
type PlayResp struct {
	X1   int32 `json:"x1"`
	Y1   int32 `json:"y1"`
	X2   int32 `json:"x2"`
	Y2   int32 `json:"y2"`
	Step int   `json:"step"` // 步数验证
}

// InviteForm contains invitor and invitee of a invitation
type InviteForm struct {
	Invitor string `json:"invitor"` // 邀请者
	Invitee string `json:"invitee"` // 受邀者
}
