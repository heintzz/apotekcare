package merchant

import (
	"context"
	"log"
)

type repositoryContract interface {
	getMerchantProfile(ctx context.Context) (merchant Merchant, err error)
	updateMerchantProfile(ctx context.Context, merchant Merchant) (err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) merchantProfile(ctx context.Context) (newMerchant MerchantResponse, err error) {	
	merchant, err := s.repo.getMerchantProfile(ctx)
	if err != nil {
		log.Println("[merchantProfile, getMerchantProfile] error : ", err)
		return
	}

	newMerchant = MerchantResponse{
		Id:       merchant.Id,
		Email:    merchant.Email,
		Name:     merchant.Name,
		ImageUrl: merchant.ImageUrl,
		City:     merchant.City,
		Address:  merchant.Address,
	}

	return newMerchant, nil
}

func (s service) editMerchant(ctx context.Context, req editMerchantRequest) (err error) {	
	if err = req.Validate(); err != nil {
		return
	}

	merchant := NewMerchant(req.Name, req.ImageUrl, req.City, req.Address)
	err = s.repo.updateMerchantProfile(ctx, merchant)
	if err != nil {
		log.Println("[editMerchant, updateMerchantProfile] error :", err)
		return
	}

	return
}