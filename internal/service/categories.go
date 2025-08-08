package service

import (
	"errors"
	"github.com/Ira11111/ProductService/internal/storage"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type CategoryProvider interface {
	Categories(c *gin.Context) ([]*api.Category, error)
	SaveCategory(c *gin.Context, categoryName string) (*api.Category, error)
	DropCategory(c *gin.Context, id int64) error
}

func (s *ServiceAPI) Categories(c *gin.Context) ([]*api.Category, error) {
	const op = "service.Categories"
	log := s.logger.With(slog.String("op", op))
	log.Info("trying to get categories")
	res, err := s.storage.Categories(c)
	if err != nil {
		log.Warn("failed to get categories")
		if errors.Is(err, storage.ErrEntityNotFound) {
			return nil, ErrFailedToFindEntity
		}
		return nil, err
	}
	return res, nil
}

func (s *ServiceAPI) CreateCategory(c *gin.Context, category *api.Category) (*api.Category, error) {
	const op = "service.Categories.CreateCategory"
	log := s.logger.With(slog.String("op", op))

	log.Info("checking token")
	isAdmin, err := s.isAdmin(c)
	if err != nil {
		return nil, err
	}
	if !isAdmin {
		return nil, ErrInvalidPermission
	}

	log.Info("trying to create category")
	newCategory, err := s.storage.SaveCategory(c, *category.Name)
	if err != nil {
		log.Warn("failed to create category")
		return nil, err
	}
	return newCategory, nil
}

func (s *ServiceAPI) DeleteCategory(c *gin.Context, id int64) error {
	const op = "service.Categories.DeleteCategory"
	log := s.logger.With(slog.String("op", op))
	log.Info("checking token")
	isAdmin, err := s.isAdmin(c)
	if err != nil {
		return err
	}
	if !isAdmin {
		return ErrInvalidPermission
	}

	log.Info("trying to delete category")
	err = s.storage.DropCategory(c, id)
	if err != nil {
		log.Warn("failed to delete category")
		if errors.Is(err, storage.ErrEntityNotFound) {
			return ErrFailedToFindEntity
		}
		return err
	}
	return nil
}
