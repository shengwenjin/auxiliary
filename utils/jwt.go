package utils

import "github.com/golang-jwt/jwt/v4"

// GetJwtToken @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @uuid: 数据载体
func GetJwtToken(secretKey string, iat, seconds int64, uuid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uuid"] = uuid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}