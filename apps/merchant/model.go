package merchant

import "time"

type Merchant struct {
	Id        int
	Email     string	
	Name 		  string
	ImageUrl 	string
	City 		  string
	Address 	string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMerchant(name, imageUrl, city, address string) Merchant {
	return Merchant{
		Name: name,
		ImageUrl: imageUrl,
		City: city,
		Address: address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}