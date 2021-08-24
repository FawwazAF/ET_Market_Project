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

// func TestExtractToken(t *testing.T) {
// 	token, _ := CreateTokenTest(1)
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	var expectation float64 = 0
// 	actual := float64(ExtractToken(c))
// 	if actual != expectation {
// 		t.Errorf("Expected %v but got %v", expectation, actual)
// 	}
// }
