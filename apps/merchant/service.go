package merchant

import (
	"context"
	"log"
)

type repositoryContract interface {
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

func (s service) editMerchant(ctx context.Context, req editMerchantRequest) (err error) {	
	if err = req.Validate(); err != nil {
		return
	}

	merchant := Merchant{
		Name: req.Name,
		ImageUrl: req.ImageUrl,
		City: req.City,
		Address: req.Address,
	} 

	err = s.repo.updateMerchantProfile(ctx, merchant)
	if err != nil {
		log.Println("[editMerchant, updateMerchantProfile] error :", err)
		return
	}

	return
}