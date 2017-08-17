package middlewares

// Session
type Session struct {
	UserID string `json:"user_id"` // 用户id
	Nick   string `json:"nick"`    // 用户昵称
}

func GetSession() {

}
