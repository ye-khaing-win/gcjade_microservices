package service

import (
	"context"
	"gcjade/services/catalogue-service/internal/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	return s.repo.Create(ctx, category)
}

func (s *CategoryService) List(ctx context.Context) ([]*domain.Category, error) {
	return s.repo.List(ctx)
}

func (s *CategoryService) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	return s.repo.FindByID(ctx, id)
}
