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

func (h handler) getDetailProductHandler(w http.ResponseWriter, r *http.Request) {
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

	resp := helper.APIResponse{
		HttpCode: http.StatusOK,
		Message: "SUCCESS",
		Payload: product,
	}

	resp.WriteJsonResponse(w)
}

func (h handler) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query().Get("query")

	products, err := h.svc.getProducts(r.Context(), queryParams)
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
		HttpCode: http.StatusOK,
		Message: "SUCCESS",
		Payload: products,
		Query: queryParams,
	}

	resp.WriteJsonResponse(w)
}