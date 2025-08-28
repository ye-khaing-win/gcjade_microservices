package repository

import (
	"context"
	"gcjade/services/catalogue-service/internal/domain"
)

type InmemRepository struct {
	categories map[string]*domain.Category
}

func NewInmemRepository() *InmemRepository {
	return &InmemRepository{
		categories: make(map[string]*domain.Category),
	}
}

func (r *InmemRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	r.categories[category.ID.Hex()] = category

	return category, nil
}

func (r *InmemRepository) List(ctx context.Context) ([]*domain.Category, error) {
	var categories []*domain.Category
	for _, category := range r.categories {
		categories = append(categories, category)
	}
	return categories, nil

}
