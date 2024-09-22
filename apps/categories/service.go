package categories

import (
	"context"
	"log"
)

type repositoryContract interface {
	addNewCategory(ctx context.Context, category Category) (err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) addCategory(ctx context.Context, req createCategoryRequest) (err error) {
	if err = req.Validate(); err != nil {
		return 
	}

	category := NewCategory(req.Name)
	err = s.repo.addNewCategory(ctx, category)
	if err != nil {
		log.Println("[addCategory, addNewCategory] error :", err)
		return
	}

	return
}