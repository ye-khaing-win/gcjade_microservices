package repository

import (
	"context"
	"errors"
	"gcjade/services/catalogue-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository struct {
	db *mongo.Database
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Create(ctx context.Context, category *domain.Category) (*domain.Category, error) {
	result, err := r.db.Collection("categories").InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)

	return category, nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]*domain.Category, error) {
	filter := bson.D{{}}
	findOptions := options.Find()
	cursor, err := r.db.Collection("categories").Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var categories []*domain.Category
	for cursor.Next(ctx) {
		var category domain.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": _id}
	result := r.db.Collection("categories").FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var category domain.Category
	if err := result.Decode(&category); err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}
