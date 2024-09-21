package utils

import (
	"fmt"
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
	Email 	string 		`json:"email"`
	Role 		string 		`json:"role"`
	Expires time.Time `json:"expires"`
}

func NewJWT(email string, role string) JWT {
	return JWT{
		Email: email,
		Role: role,
		Expires: time.Now().Add(time.Duration(expired) * time.Minute),
	}
}

func (j JWT) GenerateToken() (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": j.Email,
		"role": j.Role,
		"expires": j.Expires,
	})
	
	tokenString, err = token.SignedString([]byte(secretKey))
	return 
}

func VerifyToken(tokenString string) (token JWT, err error) {	
	jwtToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {		
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		err = fmt.Errorf("invalid token")
		return
	}  

	email := fmt.Sprintf("%v", claims["email"])
	role := fmt.Sprintf("%v", claims["role"])
	expires := fmt.Sprintf("%v", claims["expires"])

	expiresTime, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return
	}

	if time.Now().After(expiresTime) {
		err = fmt.Errorf("token expired")
		return
	}

	token = NewJWT(email, role)

	return
}
