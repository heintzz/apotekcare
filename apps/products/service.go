package products

import (
	"context"
	"log"
)

type repositoryContract interface {
	addNewProduct(ctx context.Context, product Product) (err error)
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

	return product, nil	
}