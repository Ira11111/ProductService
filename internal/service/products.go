package service

import (
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"log/slog"
	"slices"
)

type ProductProvider interface {
	Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error)
	SaveProduct(c *gin.Context, product *api.ProductCreate, sellerId int64) (*api.ProductResponse, error)
	Product(c *gin.Context, id int64) (*api.ProductResponse, error)
	DeleteProduct(c *gin.Context, id int64) error
}

func (s *ServiceAPI) Products(c *gin.Context, offset int64, limit int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (s *ServiceAPI) CreateProduct(c *gin.Context, product *api.ProductCreate) (*api.ProductResponse, error) {
	const op = "service.CreateProduct"
	log := s.logger.With(
		slog.String("op", op),
	)

	log.Info("trying to get token info")
	uid, roles, err := s.tokenInfo(c)
	if err != nil {
		log.Warn("failed to get token info")
		return nil, err
	}
	log.Info("check permission")
	if !slices.Contains(roles, sellerRole) && !slices.Contains(roles, adminRole) {
		log.Warn("permission denied")
		return nil, ErrInvalidPermission
	}

	log.Info("creating product")
	newProduct, err := s.storage.SaveProduct(c, product, uid)
	if err != nil {
		log.Warn("failed to save product", slog.String("message", err.Error()))
		return nil, ErrFailedToSaveEntity
	}

	log.Info("product created")
	return newProduct, nil
}
func (s *ServiceAPI) Product(c *gin.Context, id int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (s *ServiceAPI) DeleteProduct(c *gin.Context, productId int64, sellerId int64) error {
	return nil
}
