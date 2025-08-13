package service

import (
	"log/slog"

	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

type WarehouseProvider interface {
}
type WarehouseService struct {
	logger   *slog.Logger
	provider WarehouseProvider
}

func NewWarehouseService(logger *slog.Logger, provider WarehouseProvider) *WarehouseService {
	return &WarehouseService{
		logger:   logger,
		provider: provider,
	}
}

func (s *ServiceAPI) Warehouse(c *gin.Context, id int64) (*api.Warehouse, error) {
	return nil, nil
}
func (s *ServiceAPI) CreateWarehouse(c *gin.Context, warehouse *api.Warehouse) (*api.Warehouse, error) {
	return nil, nil
}
