package middlewares

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/context"
)

const (
	JwtUserIDKey = "usr_id"
	JwtSecret    = "yi-ptj-20170730"
)

// GenerateToken generates a new jwt token with userID
func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": time.Date(2017, 7, 30, 0, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Now().AddDate(0, 1, 0).Unix(),

		"usr_id": userID,
	})
	return token.SignedString(JwtSecret)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(JwtSecret), nil
}

// ValidateToken make it sure that a jwt token string is valid
func ValidateToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, keyFunc)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}
}

// GetUserID returns userID without any validation, so it's dangerous if you use it with no validation
func GetUserID(ctx context.Context) string {
	token, _ := jwt.Parse(ctx.GetHeader(JwtTokenHttpHeaderName), keyFunc)
	return token.Claims.(jwt.MapClaims)[JwtUserIDKey].(string)
}
