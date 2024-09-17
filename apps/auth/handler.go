package auth

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

func (h handler) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp := helper.APIResponse{
			Success:  false,
			Status: http.StatusBadRequest,
			Message: "bad request",
			Error:   err.Error(),			
		}

		resp.WriteJsonResponse(w)
		return
	}

	auth := NewAuth(req.Email, req.Password, req.Fullname)
	err = h.svc.createUser(auth)
	if err != nil {
		resp := helper.APIResponse{
			Success: false,
			Status: http.StatusBadRequest,
			Message: "internal server error",
			Error:   err.Error(),
		}
		resp.WriteJsonResponse(w)
		return
	}

	resp := helper.APIResponse{
		Success:  true,
		Status: http.StatusOK,
		Message: "registration success",
	}
	resp.WriteJsonResponse(w)
}