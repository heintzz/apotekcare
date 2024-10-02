package auth

import (
	"bytes"
	"encoding/json"
	"heintzz/apotekcare/internal/constants"
	"heintzz/apotekcare/internal/helper"
	"io"
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

func (h handler) registerHandler(w http.ResponseWriter, r *http.Request) {
    bodyBytes, err := io.ReadAll(r.Body)
    if err != nil {
        resp := helper.APIResponse{
            HttpCode: http.StatusBadRequest,
            Success:  false,
            Message:  "bad request",
            Error:    err.Error(),
        }
        resp.WriteJsonResponse(w)
        return
    }
		r.Body.Close()

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var req registerRequest
    err = json.Unmarshal(bodyBytes, &req)
    if err != nil {
        resp := helper.APIResponse{
            HttpCode: http.StatusBadRequest,
            Success:  false,
            Message:  "bad request",
            Error:    err.Error(),
        }
        resp.WriteJsonResponse(w)
        return
    }
    
    switch req.Role {
    case constants.ROLE_MERCHANT:
        var merchantReq registerRequestMerchant        
        err = json.Unmarshal(bodyBytes, &merchantReq)
        if err != nil {
            resp := helper.APIResponse{
                HttpCode: http.StatusBadRequest,
                Success:  false,
                Message:  "bad request",
                Error:    err.Error(),
            }
            resp.WriteJsonResponse(w)
            return
        }

        err = h.svc.createUser(merchantReq)
        if err != nil {
            errors, ok := helper.ErrorMapping[err.Error()]
            if !ok {
                errors = helper.ErrorGeneral
            }
            resp := helper.APIResponse{
                HttpCode:  errors.HttpCode,
                Success:   false,
                Message:   errors.ErrorMessage(),
                Error:     err.Error(),
                ErrorCode: errors.Code,
            }
            resp.WriteJsonResponse(w)
            return
        }

    case constants.ROLE_USER, "": 				
        var userReq registerRequestUser
        err = json.Unmarshal(bodyBytes, &userReq)
        if err != nil {						
            resp := helper.APIResponse{
                HttpCode: http.StatusBadRequest,
                Success:  false,
                Message:  "bad request",
                Error:    err.Error(),
            }
            resp.WriteJsonResponse(w)
            return
        }
        
        err = h.svc.createUser(userReq)
        if err != nil {
            errors, ok := helper.ErrorMapping[err.Error()]
            if !ok {
                errors = helper.ErrorGeneral
            }
            resp := helper.APIResponse{
                HttpCode:  errors.HttpCode,
                Success:   false,
                Message:   errors.ErrorMessage(),
                Error:     err.Error(),
                ErrorCode: errors.Code,
            }
            resp.WriteJsonResponse(w)
            return
        }

    default:
        resp := helper.APIResponse{
            HttpCode: http.StatusBadRequest,
            Success:  false,
            Message:  "invalid role",
            Error:    "role should be 'user' or 'merchant'",
        }
        resp.WriteJsonResponse(w)
        return
    }
   
    resp := helper.APIResponse{
        HttpCode: http.StatusOK,
        Success:  true,
        Message:  "registration success",
    }
    resp.WriteJsonResponse(w)
}

func (h handler) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest

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

	tokenString, err := h.svc.loginUser(req)
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
		Success:  true,	
		Message: "login success",
		Payload: map[string]interface{}{
			"token": tokenString,
		},
	}
	resp.WriteJsonResponse(w)
}