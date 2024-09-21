package users

import (
	"context"
	"log"
)

type repositoryContract interface {
	fetchProfile(ctx context.Context) (user User, err error)
	updateProfile(ctx context.Context, user User) (err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) getProfile(ctx context.Context) (user userProfileResponse, err error) { 
	resp, err := s.repo.fetchProfile(ctx)
	if err != nil {
		log.Println("[getProfile, fetchUserProfile] error :", err)
		return
	}

	user = userProfileResponse{
		Id: resp.Id,
		Email: resp.Email,
		FullName: resp.FullName,
		Gender: resp.Gender,
		Address: resp.Address,
		PhoneNumber: resp.PhoneNumber,
	}

	return
} 

func (s service) updateProfile(ctx context.Context, req editProfileRequest) (err error) { 
	err = req.Validate()
	if err != nil {
		return
	}

	var user = User{
		FullName: req.Fullname,
		Address: req.Address,
		Gender: req.Gender,
		PhoneNumber: req.PhoneNumber,
	}

	err = s.repo.updateProfile(ctx, user)
	if err != nil {
		log.Println("[updateProfile, updateProfile] error :", err)
		return
	}

	return
} 