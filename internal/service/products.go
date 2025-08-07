package service

import (
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ProductProvider interface {
	Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error)
	SaveProduct(c *gin.Context, product *api.ProductCreate) (*api.ProductResponse, error)
	Product(c *gin.Context, id int64) (*api.ProductResponse, error)
	DeleteProduct(c *gin.Context, id int64) error
}
type ProductService struct {
	logger   *slog.Logger
	provider ProductProvider
}

func NewProductService(logger *slog.Logger, provider ProductProvider) *ProductService {
	return &ProductService{
		logger:   logger,
		provider: provider}
}

func (p *ProductService) Products(c *gin.Context, offset int64, limit int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (p *ProductService) CreateProduct(c *gin.Context, product *api.ProductCreate, sellerId int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (p *ProductService) Product(c *gin.Context, id int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (p *ProductService) DeleteProduct(c *gin.Context, productId int64, sellerId int64) error {
	return nil
}
