package products

import "log/slog"

type ProductProvider interface{}
type ProductService struct {
	logger   *slog.Logger
	provider ProductProvider
}

func NewProductService(logger *slog.Logger, provider ProductProvider) *ProductService {
	return &ProductService{
		logger:   logger,
		provider: provider}
}

func (p *ProductService) GetProduct() error {
	return nil
}
