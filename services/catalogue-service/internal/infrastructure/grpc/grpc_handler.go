package grpc

import (
	"context"
	"errors"
	"gcjade/services/catalogue-service/internal/domain"
	"gcjade/services/catalogue-service/internal/infrastructure/repository"
	pb "gcjade/shared/proto/catalogue"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		Name:        req.GetName(),
		Description: req.GetDescription(),
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category: %v", err)
	}

	return &pb.Category{
		Id:          category.ID.Hex(),
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (h *Handler) ListCategories(ctx context.Context, req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	categories, err := h.categoryService.List(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list categories: %v", err)
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

func (h *Handler) FindCategoryByID(ctx context.Context, req *pb.FindCategoryByIDRequest) (*pb.Category, error) {
	id := req.GetId()

	category, err := h.categoryService.FindByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrCategoryNotFound):
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID.Hex(),
		Name:        category.Name,
		Description: category.Description,
	}, err

}
