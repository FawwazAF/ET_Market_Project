package middlewares

import (
	"etmarket/project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//create token with adding limit time
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func ExtractToken(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}
	return 0
}

// func ExtractToken(e echo.Context) int {
// 	if temp := e.Get("user"); temp != nil {
// 		user := e.Get("user").(*jwt.Token)
// 		if user.Valid {
// 			claims := user.Claims.(jwt.MapClaims)
// 			return int(claims["userId"].(float64))
// 		}
// 	}
// 	return 0
// }
