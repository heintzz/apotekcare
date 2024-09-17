package auth

import (
	"database/sql"
	"fmt"
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

func (s service) createUser(auth Auth) (err error) {
	existingAuth, err := s.repo.getByEmail(auth.Email)
	if err != nil && err != sql.ErrNoRows { 
		log.Println("error when trying to check if email exists", err.Error())
		return
	}

	if existingAuth.Email != "" {
		log.Println("email already registered")
		return fmt.Errorf("email already registered")
	}

	auth.Password, err = utils.Hash(auth.Password)
	if err != nil {
		log.Println("error when try to hash password with error", err.Error())
		return
	}

	err = s.repo.registerUser(auth)
	if err != nil {
		log.Println("error when try to create user with error", err.Error())
		return
	}

	return
}