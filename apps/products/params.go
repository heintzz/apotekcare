package products

import "heintzz/ecommerce/internal/helper"

type addProductRequest struct {
	Name        string	`json:"name"`
	ImageUrl    string	`json:"image_url"`
	Price       int			`json:"price"`
	Stock       int			`json:"stock"`
	Description string	`json:"description"`
	CategoryId  int			`json:"category_id"`
}

type checkoutProductRequest struct {
	ProductId int `json:"product_id"`
	Quantity 	int `json:"quantity"`
}

type categoryResponse struct {
	Id 	 int 		`json:"id"`
	Name string `json:"name"`
}

type merchantResponse struct {
	Id 	 int 		`json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type DetailProduct struct {
	Id         	int     `json:"id"`
	Name      	string  `json:"name"`
	ImageUrl    string  `json:"image"`
	Price       int			`json:"price"`
	Stock       int			`json:"stock"`
	Description string	`json:"description"`
	Category    categoryResponse `json:"category"`
	Merchant 		merchantResponse `json:"merchant"`
}

type ProductResponse struct {
	Id         	int     `json:"id"`
	Name      	string  `json:"name"`
	ImageUrl    string  `json:"image"`
	Price       int			`json:"price"`		
	Category    categoryResponse `json:"category"`
	Merchant 		merchantResponse `json:"merchant"`
}

func (req addProductRequest) Validate() (err error) {
	if err := req.ValidateName(); err != nil {
		return err
	}
	if err := req.ValidateImageUrl(); err != nil {
		return err
	}
	if err := req.ValidatePrice(); err != nil {
		return err
	}
	if err := req.ValidateStock(); err != nil {
		return err
	}
	if err := req.ValidateDescription(); err != nil {
		return err
	}
	if err := req.ValidateCategoryId(); err != nil {
		return err
	}
	return nil
}

func (req addProductRequest) ValidateName() (err error) {
	if req.Name == "" {
		return helper.ErrProductNameRequired
	}
	return nil
}

func (req addProductRequest) ValidateImageUrl() (err error) {
	if req.ImageUrl == "" {
		return helper.ErrProductImageUrlRequired
	}
	return nil
}

func (req addProductRequest) ValidatePrice() (err error) {
	if req.Price <= 0 {
		return helper.ErrProductPriceInvalid
	}
	return nil
}

func (req addProductRequest) ValidateStock() (err error) {
	if req.Stock <= 0 {
		return helper.ErrProductStockInvalid
	}
	return nil
}

func (req addProductRequest) ValidateDescription() (err error) {
	if req.Description == "" {
		return helper.ErrProductDescriptionRequired
	}
	return nil
}

func (req addProductRequest) ValidateCategoryId() (err error) {
	if req.CategoryId <= 0  {
		return helper.ErrProductCategoryIdRequired
	}
	return nil
}
type getProductRequest struct {
	Id string 
}

func (req getProductRequest) ValidateId() (err error) {
	if req.Id == "" {
		return helper.ErrProductIdRequired
	}
	return nil
}