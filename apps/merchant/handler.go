package merchant

import (
	"encoding/json"
	"heintzz/ecommerce/internal/helper"
	"net/http"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) getMerchantProfile(w http.ResponseWriter, r *http.Request) {
	merchant, err := h.svc.merchantProfile(r.Context())

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
		Payload: merchant,
	}

	resp.WriteJsonResponse(w)
}

func (h handler) editMerchantHandler(w http.ResponseWriter, r *http.Request) {
	var req editMerchantRequest

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

	err = h.svc.editMerchant(r.Context(), req)	
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
	}

	resp.WriteJsonResponse(w)
}