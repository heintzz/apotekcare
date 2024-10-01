package products

import (
	"encoding/json"
	"heintzz/ecommerce/internal/helper"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) addProductHandler(w http.ResponseWriter, r *http.Request) {
	var req addProductRequest
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := helper.APIResponse{
			HttpCode: http.StatusBadRequest,
			Success:  false,
			Message: "bad request",
			Error:   err.Error(),			
		}

		resp.WriteJsonResponse(w)
		return
	}

	err = h.svc.addProduct(r.Context(), req)
	if err != nil {
		errors, ok := helper.ErrorMapping[err.Error()]
		if !ok {
			errors = helper.ErrorGeneral
		}
		resp := helper.APIResponse{
			HttpCode: errors.HttpCode,
			Success: false,
			Message: errors.ErrorMessage(),
			Error:   err.Error(),
			ErrorCode: errors.Code,
		}

		resp.WriteJsonResponse(w)
		return
	}

	resp := helper.APIResponse{
		HttpCode: http.StatusCreated,
		Message: "SUCCESS",
	}

	resp.WriteJsonResponse(w)
}

func (h handler) getProductHandler(w http.ResponseWriter, r *http.Request) {
	var req getProductRequest
	req.Id = chi.URLParam(r, "id")		
	
	product, err := h.svc.getProductData(r.Context(), req)

	if err != nil {
		errors, ok := helper.ErrorMapping[err.Error()]
		if !ok {
			errors = helper.ErrorGeneral
		}
		resp := helper.APIResponse{
			HttpCode: errors.HttpCode,
			Success: false,
			Message: errors.ErrorMessage(),
			Error:   err.Error(),
			ErrorCode: errors.Code,
		}

		resp.WriteJsonResponse(w)
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

	resp := helper.APIResponse{
		HttpCode: http.StatusOK,
		Message: "SUCCESS",
		Payload: detailProductResponse{
			Id: product.Id,
			Name: product.Name,
			Price: product.Price,
			Image: product.ImageUrl,
			Stock: product.Stock,
			Description: product.Description,
			Merchant: merchant,
			Category: category,
		},
	}

	resp.WriteJsonResponse(w)

}