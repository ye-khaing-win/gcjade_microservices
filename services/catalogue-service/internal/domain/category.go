package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	ID          primitive.ObjectID
	Name        string
	Description string
	ImageURL    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
}

type CategoryService interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
}
