package auth

import (
	"fmt"
	"heintzz/apotekcare/internal/helper"
	"regexp"
)

type RegisterRequest interface {
	Validate() error
	GetEmail() string
	GetPassword() string
	GetRole() string
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role 		 string `json:"role"`
}

type registerRequestUser struct {
	registerRequest
	Fullname string `json:"fullname"`
}

type registerRequestMerchant struct {
	registerRequest
	Name string 		`json:"merchant_name"`
	Address string  `json:"merchant_address"`
}

// GLOBAL
func (req registerRequest) GetEmail() string {
	return req.Email
}
func (req registerRequest) GetPassword() string {
	return req.Password
}
func (req registerRequest) GetRole() string {
	return req.Role
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

// USER VALIDATION
func (req registerRequestUser) ValidateFullname() error {
	if req.Fullname == "" {
		return helper.ErrFullnameRequired
	} 
	return nil
}

func (req registerRequestUser) Validate() error {
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

// MERCHANT VALIDATION
func (req registerRequestMerchant) ValidateMerchantName() error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func (req registerRequestMerchant) ValidateMerchantAddress() error {
	if req.Address == "" {
		return fmt.Errorf("address is required")
	}
	return nil
}

func (req registerRequestMerchant) Validate() error {
	if err := req.ValidateEmail(); err != nil {
		return err
	}
	if err := req.ValidatePassword(); err != nil {
		return err
	}
	if err := req.ValidateMerchantName(); err != nil {
		return err
	}
	if err := req.ValidateMerchantAddress(); err != nil {
		return err
	}
	return nil
}

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (req loginRequest) Validate() error {
	if err := req.ValidateEmail(); err != nil {
		return err
	}
	if err := req.ValidatePassword(); err != nil {
		return err
	}
	return nil
}

func (req loginRequest) ValidateEmail() error {
	if req.Email == "" {
		return helper.ErrEmailRequired
	}
	
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if match, _ := regexp.MatchString(emailRegex, req.Email); !match {
		return helper.ErrEmailInvalid
	}

	return nil
}

func (req loginRequest) ValidatePassword() error {
	if req.Password == "" {
		return helper.ErrPasswordRequired
	}
	if len(req.Password) < 6 {
		return helper.ErrPasswordInvalidLength
	}
	return nil
}
