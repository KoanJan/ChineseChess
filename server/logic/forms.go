package logic

// PlayForm contains the all parameters that the Play API needs
type PlayForm struct {
	X1      int32  `json:"x_1"`
	Y1      int32  `json:"y_1"`
	X2      int32  `json:"x_2"`
	Y2      int32  `json:"y_2"`
	BoardID string `json:"board_id"`
	UserID  string `json:"user_id"`
}
