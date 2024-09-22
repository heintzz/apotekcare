package categories

import "heintzz/ecommerce/internal/helper"

type createCategoryRequest struct {
	Name string `json:"name"`
}

func (req createCategoryRequest) Validate() (err error) {
	if name := req.Name; name == "" {
		return helper.ErrCategoryNameRequired
	}
	return nil
}