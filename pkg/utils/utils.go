package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var JwtSecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username, password string) (string, error) {
	var notTime = time.Now()
	var expireTime = notTime.Add(24 * time.Hour)
	var claims = Claims{
		Id:       id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "web",
		},
	}
	var tokenClaims = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var token, err = tokenClaims.SignedString(JwtSecret)
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	var tokenClaims, err = jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		var claims, ok = tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
func GetIdByToken(c *gin.Context) uint {
	claim, _ := ParseToken(c.GetHeader("Authorization"))
	return claim.Id
}
