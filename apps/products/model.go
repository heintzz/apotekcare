package products

import "time"

type Product struct {
	Id          int
	Name        string
	CategoryId  int
	MerchantId  int
	Stock       int
	Price       int
	ImageUrl    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name, image_url, description string, category_id, merchant_id, stock, price int) Product {
	return Product{
		Name: name,
		CategoryId: category_id,
		MerchantId: merchant_id,
		Stock: stock,
		Price: price,
		ImageUrl: image_url,
		Description: description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}