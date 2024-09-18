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
	ErrFullnameRequired			 = errors.New("fullname is required")
	ErrPasswordInvalidLength = errors.New("password must be at least 6 characters long")
	ErrEmailAlreadyUsed      = errors.New("email already used")
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
	ErrorGeneral         			 = NewError("internal server error", "unknown error", "99999", http.StatusInternalServerError)
	ErrorEmailRequired         = NewError(ErrBadRequest.Error(), ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrBadRequest.Error(), ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrBadRequest.Error(), ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorFullnameRequired      = NewError(ErrBadRequest.Error(), ErrFullnameRequired.Error(), "40004", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrBadRequest.Error(), ErrPasswordInvalidLength.Error(), "40005", http.StatusBadRequest)
	ErrorEmailAlreadyUsed 		 = NewError("duplicate entry", ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
)

var (
	ErrorMapping = map[string]Error{		
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrFullnameRequired.Error(): 			ErrorFullnameRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
	}
)

