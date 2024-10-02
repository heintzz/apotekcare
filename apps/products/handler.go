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
		Success: true,
		Message: "SUCCESS",
	}

	resp.WriteJsonResponse(w)
}

func (h handler) getDetailProductHandler(w http.ResponseWriter, r *http.Request) {
	var req getProductRequest
	req.Id = chi.URLParam(r, "id")		
	
	product, err := h.svc.getProduct(r.Context(), req)

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
		Success: true,
		Message: "get product detail success",
		Payload: product,
	}

	resp.WriteJsonResponse(w)
}

func (h handler) getDetailProductByMerchant(w http.ResponseWriter, r *http.Request) {
	var req getProductRequest
	req.Id = chi.URLParam(r, "id")		
	
	product, err := h.svc.merchantProduct(r.Context(), req)

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
		Success: true,
		Message: "get product detail success",
		Payload: product,
	}

	resp.WriteJsonResponse(w)
}

func (h handler) getProductsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query().Get("query")

	products, err := h.svc.products(r.Context(), queryParams)
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
		Success: true,
		Message: "get product success",
		Payload: products,
		Query: queryParams,
	}

	resp.WriteJsonResponse(w)
}


func (h handler) getProductsByMerchantHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query().Get("query")

	products, err := h.svc.merchantProducts(r.Context(), queryParams)
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
		Success: true,
		Message: "get product success",
		Payload: products,
		Query: queryParams,	
	}

	resp.WriteJsonResponse(w)
}

func (h handler) checkoutProductHandler(w http.ResponseWriter, r *http.Request) {
	var req checkoutProductRequest

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


	err = h.svc.checkoutProduct(r.Context(), req)
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
		Success: true,
		Message: "checkout product success",		
	}

	resp.WriteJsonResponse(w)
}