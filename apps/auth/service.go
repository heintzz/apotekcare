package auth

import (
	"database/sql"
	"fmt"
	"heintzz/apotekcare/internal/constants"
	"heintzz/apotekcare/internal/helper"
	"heintzz/apotekcare/internal/utils"
	"log"
)

type repositoryContract interface {
	registerUser(auth Auth) (err error)
	getByEmail(email string) (auth Auth, err error)
	insertToUsersTable(email, fullname string) (err error)
	insertToMerchantsTable(email, name, address string) (err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) createUser(req RegisterRequest) (err error) {
	err = req.Validate()
	if err != nil {
		return
	}

	existingAuth, err := s.repo.getByEmail(req.GetEmail())
	if err != nil && err != sql.ErrNoRows { 
		log.Println("error when trying to check if email exists", err.Error())
		return
	}

	if existingAuth.Email != "" {
		log.Println("error registering user with error email already used")
		return helper.ErrEmailAlreadyUsed
	}

	hashedPassword, err := utils.Hash(req.GetPassword())
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	var role string
	if req.GetRole() == "" {
		role = "user"
	} else {
		role = req.GetRole()
	}

	auth := NewAuth(req.GetEmail(), hashedPassword, role)

	err = s.repo.registerUser(auth)
	if err != nil {
		log.Println("error when try to create user with error", err.Error())
		return
	}

	if role == constants.ROLE_USER {
		request := req.(registerRequestUser)
		return s.repo.insertToUsersTable(request.Email, request.Fullname)
	} else if role == constants.ROLE_MERCHANT {
		request := req.(registerRequestMerchant)
		return s.repo.insertToMerchantsTable(request.Email, request.Name, request.Address)
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

	tokenJWT := utils.NewJWT(auth.Email, auth.Role)
	token, err = tokenJWT.GenerateToken()
	if err != nil {
		log.Println("error when try to GenerateToken with error", err.Error())
		return
	}
	
	return
}