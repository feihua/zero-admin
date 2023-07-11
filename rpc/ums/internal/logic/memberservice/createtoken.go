package memberservicelogic

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//创建token
func createJwtToken(secretKey, name, mobile string, seconds, memberId int64) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["memberId"] = memberId
	claims["memberName"] = name
	claims["mobile"] = mobile
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
