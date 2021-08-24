package middlewares

import (
	"etmarket/project/constants"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateTokenTest(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func TestCreateToken(t *testing.T) {
	expectation, _ := CreateTokenTest(1)
	actual, _ := CreateToken(1)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
