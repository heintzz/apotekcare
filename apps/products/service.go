package products

import (
	"context"
	"fmt"
	"log"
)

type repositoryContract interface {
	addNewProduct(ctx context.Context, product Product) (err error)
	getProducts(ctx context.Context, queryParams string) (products []ProductResponse, err error)
	getProductsByMerchant(ctx context.Context, queryParams string) (products []ProductResponse, err error)
	getDetailProduct(ctx context.Context, productId string) (product DetailProduct, err error)
	getDetailProductByMerchant(ctx context.Context, productId string) (product DetailProduct, err error)
	getProductStock(ctx context.Context, req checkoutProductRequest) (stock int, err error) 
	checkoutProduct(ctx context.Context, req checkoutProductRequest) (err error) 
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
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

func (s service) products(ctx context.Context, queryParams string) (productsResponse []ProductResponse, err error) {
	products, err := s.repo.getProducts(ctx, queryParams)
	if err != nil {
		return
	}

	for _, product := range products {		
		merchant := merchantResponse{
			Id: product.Merchant.Id,
			Name: product.Merchant.Name,
			City: product.Merchant.City,
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

func (s service) merchantProducts(ctx context.Context, queryParams string) (productsResponse []ProductResponse, err error) {
	products, err := s.repo.getProductsByMerchant(ctx, queryParams)
	if err != nil {
		return
	}

	if len(products) == 0 {
		return []ProductResponse{}, nil
	}

	for _, product := range products {		
		merchant := merchantResponse{
			Id: product.Merchant.Id,
			Name: product.Merchant.Name,
			City: product.Merchant.City,
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

func (s service) getProduct(ctx context.Context, req getProductRequest) (product DetailProduct, err error) {
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

func (s service) merchantProduct(ctx context.Context, req getProductRequest) (product DetailProduct, err error) {
	if err = req.ValidateId(); err != nil {
		return
	}
	
	product, err = s.repo.getDetailProductByMerchant(ctx, req.Id)
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

func (s service) checkoutProduct(ctx context.Context, req checkoutProductRequest) (err error) {
	stock, err := s.repo.getProductStock(ctx, req) 
	if err != nil {
		return
	}

	if stock < req.Quantity {
		return fmt.Errorf("insufficient stock: available stock is %d, but %d requested", stock, req.Quantity)
	}

	err = s.repo.checkoutProduct(ctx, req)
	if err != nil {
		return
	}

	return
}