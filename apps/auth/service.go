package auth

import (
	"database/sql"
	"fmt"
	"heintzz/ecommerce/internal/helper"
	"heintzz/ecommerce/internal/utils"
	"log"
)

type repositoryContract interface {
	registerUser(auth Auth) (err error)
	getByEmail(email string) (auth Auth, err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) createUser(req registerRequest) (err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	existingAuth, err := s.repo.getByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows { 
		log.Println("error when trying to check if email exists", err.Error())
		return
	}

	if existingAuth.Email != "" {
		log.Println("error registering user with error email already used")
		return helper.ErrEmailAlreadyUsed
	}

	req.Password, err = utils.Hash(req.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	auth := NewAuth(req.Email, req.Password, req.Fullname)

	err = s.repo.registerUser(auth)
	if err != nil {
		log.Println("error when try to create user with error", err.Error())
		return
	}

	return
}	

func (s service) loginUser(req loginRequest) (token string, err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	auth, err := s.repo.getByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows { 
		log.Println("error when trying to check if email exists", err.Error())
		return
	}

	if auth.Email == "" {
		log.Println("user is not found")
		err = fmt.Errorf("user is not found")
		return 
	}

	err = utils.Verify(req.Password, auth.Password)
	if err != nil {
		log.Println("error when try to verify password with error", err.Error())
		err = fmt.Errorf("invalid credentials")
		return
	}

	tokenJWT := utils.NewJWT(auth.Id, auth.Role)
	token, err = tokenJWT.GenerateToken()
	if err != nil {
		log.Println("error when try to GenerateToken with error", err.Error())
		return
	}
	return
}