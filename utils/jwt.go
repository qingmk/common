package utils

import "github.com/dgrijalva/jwt-go"

func GetJwtToken(secretKey string, iat, seconds, userId int64, userName string, roleIds []int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["userName"] = userName
	claims["roleIds"] = roleIds
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GetJwtTokenV2(secretKey string, iat, seconds, userId int64, nameSpaceId int64, userName string, roleIds []int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["nameSpaceId"] = nameSpaceId
	claims["userId"] = userId
	claims["userName"] = userName
	claims["roleIds"] = roleIds
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
