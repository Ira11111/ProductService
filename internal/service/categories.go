package service

import (
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type CategoryProvider interface {
	Categories() ([]*api.Category, error)
	SaveCategory(categoryName string) (*api.Category, error)
	DropCategory(id int64) error
}

type CategoryService struct {
	logger   *slog.Logger
	provider CategoryProvider
}

func NewCategoryService(logger *slog.Logger, provider CategoryProvider) *CategoryService {
	return &CategoryService{
		logger:   logger,
		provider: provider,
	}
}

func (s *CategoryService) Categories(c *gin.Context) ([]*api.Category, error) {
	return nil, nil
}

func (s *CategoryService) CreateCategory(c *gin.Context) (*api.Category, error) {
	return nil, nil
}

func (s *CategoryService) DeleteCategory(c *gin.Context, id int) error {
	return nil
}
