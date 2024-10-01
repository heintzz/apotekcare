package merchant

import "heintzz/ecommerce/internal/helper"

type editMerchantRequest struct {
	Name     string `json:"merchant_name"`
	ImageUrl string `json:"image_url"`
	City     string `json:"merchant_city"`
	Address  string `json:"merchant_address"`
}

type MerchantResponse struct {
	Id 				int 	 `json:"id"`
	Email 		string `json:"email"`
	Name 			string `json:"name"`
	ImageUrl 	string `json:"image_url"`
	City 			string `json:"city"`
	Address 	string `json:"address"`
}

func (req editMerchantRequest) Validate() (err error) {
	if err := req.ValidateName(); err != nil {
		return err
	}
	if err := req.ValidateCity(); err != nil {
		return err
	}
	if err := req.ValidateAddress(); err != nil {
		return err
	}
	if err := req.ValidateImageUrl(); err != nil {
		return err
	}
	return nil
}

func (req editMerchantRequest) ValidateName() (err error) {
	if req.Name == "" {
		return helper.ErrMerchantNameRequired
	}
	if len(req.Name) < 3 {
		return helper.ErrMerchantNameInvalid
	}
	return nil
}

func (req editMerchantRequest) ValidateCity() (err error) {
	if req.City == "" {
		return helper.ErrMerchantCityRequired
	}
	return nil
}

func (req editMerchantRequest) ValidateAddress() (err error) {
	if req.Address == "" {
		return helper.ErrMerchantAddressRequired
	}
	return nil
}

func (req editMerchantRequest) ValidateImageUrl() (err error) {
	if req.ImageUrl == "" {
		return helper.ErrMerchantImageUrlRequired
	}
	return nil
}