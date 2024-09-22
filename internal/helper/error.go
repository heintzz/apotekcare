package helper

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
	ErrBadRequest 		 = errors.New("bad request")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must be at least 6 characters long")
	ErrEmailAlreadyUsed      = errors.New("email already used")

	ErrFullnameRequired			   = errors.New("fullname is required")
	ErrFullnameInvalid			   = errors.New("fullname must be at least 3 characters long")
	ErrUserGenderRequired			 = errors.New("gender is required")
	ErrUserAddressRequired	   = errors.New("address is required")
	ErrUserAddressInvalid	     = errors.New("address must be at least 3 characters long")
	ErrUserPhoneNumberRequired = errors.New("phone number is required")

	ErrMerchantNameRequired	    = errors.New("merchant name is required")
	ErrMerchantNameInvalid      = errors.New("merchant name is invalid")
	ErrMerchantAddressRequired  = errors.New("merchant address is required")
	ErrMerchantCityRequired	    = errors.New("merchant city is required")
	ErrMerchantImageUrlRequired	= errors.New("merchant image is required")

	ErrCategoryNameRequired	= errors.New("category name is required")
)

type Error struct {
	Message  string
	Error		 string
	Code     string
	HttpCode int
}

func (e Error) ErrorMessage() string {
	return e.Message
}

func NewError(msg string, err string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Error:    err,
		Code:     code,
		HttpCode: httpCode,
	}
}

var (	
	ErrorEmailRequired         = NewError(ErrBadRequest.Error(), ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrBadRequest.Error(), ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrBadRequest.Error(), ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrBadRequest.Error(), ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	
	ErrorFullnameRequired         = NewError(ErrBadRequest.Error(), ErrFullnameRequired.Error(), "40005", http.StatusBadRequest)
	ErrorFullnameInvalid          = NewError(ErrBadRequest.Error(), ErrFullnameInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorUserGenderRequired   		= NewError(ErrBadRequest.Error(), ErrUserGenderRequired.Error(), "40007", http.StatusBadRequest)
	ErrorUserAddressRequired   		= NewError(ErrBadRequest.Error(), ErrUserAddressRequired.Error(), "40008", http.StatusBadRequest)
	ErrorUserAddressInvalid    		= NewError(ErrBadRequest.Error(), ErrUserAddressInvalid.Error(), "40009", http.StatusBadRequest)
	ErrorUserPhoneNumberRequired  = NewError(ErrBadRequest.Error(), ErrUserPhoneNumberRequired.Error(), "40010", http.StatusBadRequest)	
	
	ErrorMerchantNameRequired  		= NewError(ErrBadRequest.Error(), ErrMerchantNameRequired.Error(), "40011", http.StatusBadRequest)
	ErrorMerchantNameInvalid   		= NewError(ErrBadRequest.Error(), ErrMerchantNameInvalid.Error(), "40012", http.StatusBadRequest)
	ErrorMerchantAddressRequired  = NewError(ErrBadRequest.Error(), ErrMerchantAddressRequired.Error(), "40013", http.StatusBadRequest)
	ErrorMerchantCityRequired  		= NewError(ErrBadRequest.Error(), ErrMerchantCityRequired.Error(), "40014", http.StatusBadRequest)
	ErrorMerchantImageUrlRequired	= NewError(ErrBadRequest.Error(), ErrMerchantImageUrlRequired.Error(), "40015", http.StatusBadRequest)

	ErrorCategoryNameRequired = NewError(ErrBadRequest.Error(), ErrCategoryNameRequired.Error(), "40016", http.StatusBadRequest)
	
	ErrorEmailAlreadyUsed = NewError("duplicate entry", ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorGeneral          = NewError("internal server error", "unknown error", "99999", http.StatusInternalServerError)
)

var (
	ErrorMapping = map[string]Error{		
		ErrEmailRequired.Error():         		ErrorEmailRequired,
		ErrEmailInvalid.Error():          		ErrorEmailInvalid,
		ErrPasswordRequired.Error():      		ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): 		ErrorPasswordInvalidLength,
		ErrEmailAlreadyUsed.Error():      		ErrorEmailAlreadyUsed,
		ErrFullnameRequired.Error(): 					ErrorFullnameRequired,
		ErrFullnameInvalid.Error(): 					ErrorFullnameInvalid,
		ErrUserGenderRequired.Error():				ErrorUserGenderRequired,
		ErrUserAddressRequired.Error():				ErrorUserAddressRequired,
		ErrUserAddressInvalid.Error():				ErrorEmailAlreadyUsed,
		ErrUserPhoneNumberRequired.Error():		ErrorUserPhoneNumberRequired,
		ErrMerchantNameRequired.Error():			ErrorMerchantNameRequired,
		ErrMerchantNameInvalid.Error():	 		  ErrorMerchantNameInvalid,
		ErrMerchantAddressRequired.Error():	  ErrorMerchantAddressRequired,
		ErrMerchantCityRequired.Error():			ErrorMerchantCityRequired,
		ErrMerchantImageUrlRequired.Error():	ErrorMerchantImageUrlRequired,
		ErrCategoryNameRequired.Error(): 			ErrorCategoryNameRequired,
	}
)

