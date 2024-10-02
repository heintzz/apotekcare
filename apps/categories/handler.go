package categories

import (
	"encoding/json"
	"heintzz/apotekcare/internal/helper"
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

func (h handler) addCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req createCategoryRequest

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

	err = h.svc.addCategory(r.Context(), req)
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

func (h handler) getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := h.svc.categories(r.Context())
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

	var data []categoryResponse

	for _, category := range categories {
		data = append(data,  categoryResponse{
			Id: category.Id,
			Name: category.Name,
		})
	}

	resp := helper.APIResponse{
		HttpCode: http.StatusOK,
		Message: "get categories success",
		Payload: data,
	}

	resp.WriteJsonResponse(w)
}