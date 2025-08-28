package grpc

import (
	"context"
	"gcjade/services/catalogue-service/internal/domain"
	pb "gcjade/shared/proto/catalogue"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Handler struct {
	pb.UnimplementedCatalogueServiceServer
	categoryService domain.CategoryService
}

func NewHandler(server *grpc.Server, categoryService domain.CategoryService) *Handler {
	handler := &Handler{
		categoryService: categoryService,
	}

	pb.RegisterCatalogueServiceServer(server, handler)
	return handler
}

func (h *Handler) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := h.categoryService.Create(ctx, &domain.Category{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("Category created: %+v", category)

	return &pb.Category{
		Id:          category.ID.Hex(),
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (h *Handler) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	categories, err := h.categoryService.List(ctx)
	if err != nil {
		return nil, err
	}

	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID.Hex(),
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.ListCategoriesResponse{
		Categories: categoryList,
	}, nil
}
