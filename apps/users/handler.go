package users

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

func (h handler) getProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, err := h.svc.getProfile(r.Context())
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
		Payload: user,
	 }

	 resp.WriteJsonResponse(w)
}

func (h handler) updateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var req editProfileRequest

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

	err = h.svc.updateProfile(r.Context(), req)
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