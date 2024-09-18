package auth

import (
	"heintzz/ecommerce/internal/helper"
	"regexp"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (req registerRequest) Validate() error {
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

func (req registerRequest) ValidateEmail() error {
	if req.Email == "" {
		return helper.ErrEmailRequired
	}
	
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(emailRegex, req.Email); !match {
		return helper.ErrEmailInvalid
	}

	return nil
}

func (req registerRequest) ValidatePassword() error {
	if req.Password == "" {
		return helper.ErrPasswordRequired
	}
	if len(req.Password) < 6 {
		return helper.ErrPasswordInvalidLength
	}
	return nil
}

func (req registerRequest) ValidateFullname() error {
	if req.Fullname == "" {
		return helper.ErrFullnameRequired
	}
	return nil
}