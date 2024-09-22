package categories

import "time"

type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) Category {
	return Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}