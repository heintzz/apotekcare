package users

import "errors"

type editProfileRequest struct {
	Fullname    string `json:"fullname"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type userProfileResponse struct {
	Id         	int    
	Email      	string 
	FullName   	string 
	Gender     	string 
	Address    	string 
	PhoneNumber string 
}


func (req editProfileRequest) ValidateFullname() error {
	if req.Fullname == "" {
		return errors.New("fullname is required")
	}
	return nil
}

func (req editProfileRequest) ValidateGender() error {
	if req.Gender == "" {
		return errors.New("gender is required")
	}
	return nil
}

func (req editProfileRequest) ValidateAddress() error {
	if req.Address == "" {
		return errors.New("address is required")
	}
	return nil
}

func (req editProfileRequest) ValaidatePhoneNumber() error {
	if req.PhoneNumber == "" {
		return errors.New("phone number is required")
	}
	return nil
}

func (req editProfileRequest) Validate() error {
	if err := req.ValidateFullname(); err != nil {
		return err
	}
	if err := req.ValidateGender(); err != nil {
		return err
	}
	if err := req.ValidateAddress(); err != nil {
		return err
	}
	if err := req.ValaidatePhoneNumber(); err != nil {
		return err
	}	
	return nil
}