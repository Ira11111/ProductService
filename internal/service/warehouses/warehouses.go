package warehouses

import "log/slog"

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
