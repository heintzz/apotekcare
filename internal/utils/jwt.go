package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey string
var expired int

func InitToken(secret string, expiredToken int) {
	secretKey = secret
	expired = expiredToken
}

type JWT struct {
	Id 			int 			`json:"id"`
	Role 		string 		`json:"role"`
	Expires time.Time `json:"expires"`
}

func NewJWT(id int, role string) JWT {
	return JWT{
		Id: id,
		Role: role,
		Expires: time.Now().Add(time.Duration(expired) * time.Minute),
	}
}

func (j JWT) GenerateToken() (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id": j.Id,
		"Role": j.Role,
		"Expires": j.Expires,
	})
	
	tokenString, err = token.SignedString([]byte(secretKey))
	return 
}


