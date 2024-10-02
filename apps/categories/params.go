package categories

import "heintzz/apotekcare/internal/helper"

type createCategoryRequest struct {
	Name string `json:"name"`
}

type categoryResponse struct {
	Id 	 int `json:"id"`
	Name string `json:"name"`
}

func (req createCategoryRequest) Validate() (err error) {
	if name := req.Name; name == "" {
		return helper.ErrCategoryNameRequired
	}
	return nil
}