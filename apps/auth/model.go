package auth

import "time"

type Auth struct {
	Id        int
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdateAt 	time.Time
}

func NewAuth(email, password, role string) Auth {
	return Auth{
		Email: email,
		Password: password,		
		Role: role,
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}