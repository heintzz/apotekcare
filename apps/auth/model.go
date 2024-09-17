package auth

import "time"

type Auth struct {
	Id        int
	Email     string
	Password  string
	Role      string
	Fullname 	string
	CreatedAt time.Time
	UpdateAt 	time.Time
}

func NewAuth(email, password, fullname string) Auth {
	return Auth{
		Email: email,
		Password: password,
		Fullname: fullname,
		Role: "user",
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}