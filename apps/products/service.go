package products

import (
	"context"
	"log"
)

type repositoryContract interface {
	addNewProduct(ctx context.Context, product Product) (err error)
	getProductsLogic(ctx context.Context, queryParams string) (products []ProductResponse, err error)
	getDetailProduct(ctx context.Context, productId string) (product DetailProduct, err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) getProducts(ctx context.Context, queryParams string) (productsResponse []ProductResponse, err error) {
	products, err := s.repo.getProductsLogic(ctx, queryParams)
	if err != nil {
		return
	}

	for _, product := range products {		
		merchant := merchantResponse{
			Id: product.Merchant.Id,
			Name: product.Merchant.Name,
			// City: product.Merchant.City,
		}

		category := categoryResponse{
			Id: product.Category.Id,
			Name: product.Category.Name,
		}

		product := ProductResponse{
			Id: product.Id,
			Name: product.Name,
			Price: product.Price,
			ImageUrl: product.ImageUrl,			
			Merchant: merchant,
			Category: category,
		}

		productsResponse = append(productsResponse, product)
	}

	return
}

func (s service) addProduct(ctx context.Context, req addProductRequest) (err error) {
	if err = req.Validate(); err != nil {
		return 
	}

	product := NewProduct(
		req.Name, req.ImageUrl, req.Description, 
		req.CategoryId, 0, req.Stock, req.Price,
	)
	err = s.repo.addNewProduct(ctx, product)
	if err != nil {
		log.Println("[addProduct, addNewProduct] error : ", err)
		return
	}
	
	return
}

func (s service) getProductData(ctx context.Context, req getProductRequest) (product DetailProduct, err error) {
	if err = req.ValidateId(); err != nil {
		return
	}
	
	product, err = s.repo.getDetailProduct(ctx, req.Id)
	if err != nil {
		log.Println("[getProductData, getDetailProduct] error : ", err)
		return
	}

	merchant := merchantResponse{
		Id: product.Merchant.Id,
		Name: product.Merchant.Name,
		City: product.Merchant.City,
	}

	category := categoryResponse{
		Id: product.Category.Id,
		Name: product.Category.Name,
	}

	product = DetailProduct{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		ImageUrl: product.ImageUrl,
		Stock: product.Stock,
		Description: product.Description,
		Merchant: merchant,
		Category: category,
	}

	return 
}