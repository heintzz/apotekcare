package auth

import (
	"errors"
	"regexp"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (req RegisterRequest) Validate() error {
	if err := req.ValidateEmail(); err != nil {
		return err
	}
	if err := req.ValidatePassword(); err != nil {
		return err
	}
	if err := req.ValidateFullname(); err != nil {
		return err
	}
	return nil
}

func (req RegisterRequest) ValidateEmail() error {
	if req.Email == "" {
		return errors.New("email is required")
	}
	
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(emailRegex, req.Email); !match {
		return errors.New("invalid email format")
	}

	return nil
}

func (req RegisterRequest) ValidatePassword() error {
	if req.Password == "" {
		return errors.New("password is required")
	}
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func (req RegisterRequest) ValidateFullname() error {
	if req.Fullname == "" {
		return errors.New("fullname is required")
	}
	return nil
}