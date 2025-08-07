package postgres

import (
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

func (s *Storage) Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error) {
	return nil, nil
}
func (s *Storage) SaveProduct(c *gin.Context, product *api.ProductCreate) (*api.ProductResponse, error) {
	return nil, nil
}
func (s *Storage) Product(c *gin.Context, id int64) (*api.ProductResponse, error) {
	return nil, nil
}
func (s *Storage) DeleteProduct(c *gin.Context, id int64) error {
	return nil
}
