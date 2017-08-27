package utils

import "fmt"

// Session
type SessionField = string

const (
	// key
	SessionKeyF = "session.%s"

	// field
	SessionFieldNick   SessionField = "nick"
	SessionFieldStatus SessionField = "status"
)

// SessionKey returns a session key of redis
func SessionKey(userID string) string {
	return fmt.Sprintf(SessionKeyF, userID)
}
