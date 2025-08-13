package service

import (
	"log/slog"

	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

type SellerProvider interface{}
type SellerService struct {
	logger   *slog.Logger
	provider SellerProvider
}

func NewSellerService(logger *slog.Logger, provider SellerProvider) *SellerService {
	return &SellerService{
		logger:   logger,
		provider: provider}
}

func (s *ServiceAPI) CreateSeller(c *gin.Context, seller *api.SellerFull) (*api.SellerFull, error) {
	return nil, nil
}
func (s *ServiceAPI) Seller(c *gin.Context, id int64) (*api.SellerFull, error) {
	return nil, nil
}
func (s *ServiceAPI) DeleteSeller(c *gin.Context, id int64) error {
	return nil
}
func (s *ServiceAPI) EditSeller(c *gin.Context, full *api.SellerFull) (*api.SellerFull, error) {
	return nil, nil
}
