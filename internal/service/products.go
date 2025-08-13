package service

import (
	"log/slog"
	"slices"

	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

type ProductProvider interface {
	SaveProduct(c *gin.Context, product *api.ProductCreate, userId int64) (*api.ProductResponse, error)
	Product(c *gin.Context, id int64) (*api.ProductResponse, error)
	DeleteProduct(c *gin.Context, id int64) error
	UpdateProduct(c *gin.Context, create *api.ProductCreate) (*api.ProductCreate, error)
}

type ProductListProvider interface {
	Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error)
	ProductsByCategory(c *gin.Context, id int64, offset int64, limit int64) (*[]api.ProductResponse, error)
	ProductsByWarehouse(c gin.Context, id int64, offset int64, limit int64) (*api.ProductResponse, error)
}

func (s *ServiceAPI) Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error) {
	const op = "products.Products"
	logger := s.logger.With(slog.String("op", op))
	logger.Info("Trying to fetch products")

	products, err := s.storage.Products(c, offset, limit)
	if err != nil {
		logger.Warn("Failed to get product list", slog.String("err", err.Error()))
		return nil, ErrFailedToFindEntity
	}
	logger.Info("products are fetched successfully")
	return products, nil
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
func (s *ServiceAPI) DeleteProduct(c *gin.Context, productId int64) error {
	//1. проверить админа и роль продавца
	// если админ то все ок
	// надо булеть по id пользователя проверить что он продавец этого товара(скорее всего клиент для grpc)
	return nil
}

func (s *ServiceAPI) EditProduct(c *gin.Context, create *api.ProductCreate) (*api.ProductResponse, error) {
	//1. проверить админа и роль продавца
	// если админ то все ок
	// надо булеть по id пользователя проверить что он продавец этого товара и ему нельзя менять продавца
	// для этого создать функцию IsUserSeller service
	// в storage функция котрая находит продавца по id
	return nil, nil
}

func (s *ServiceAPI) ProductsCategory(c *gin.Context, id int64, offset int64, limit int64) (*[]api.ProductResponse, error) {
	return nil, nil
}
func (s *ServiceAPI) ProductsWarehouse(c gin.Context, id int64, offset int64, limit int64) (*api.ProductResponse, error) {
	return nil, nil
}
