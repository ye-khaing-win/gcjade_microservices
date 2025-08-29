package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	ImageURL    string             `bson:"image_url,omitempty" json:"image_url,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	FindByID(ctx context.Context, id string) (*Category, error)
}

type CategoryService interface {
	Create(ctx context.Context, category *Category) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	FindByID(ctx context.Context, id string) (*Category, error)
}
