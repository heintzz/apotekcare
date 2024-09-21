package users

import "time"

type User struct {
	Id          int
	Email       string
	FullName    string
	Address     string
	Gender      string
	PhoneNumber string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewUser(fullname, address, gender, phoneNumber string) User {
	return User{
		FullName: fullname,
		Address: address,
		Gender: gender,
		PhoneNumber: phoneNumber,
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}
}