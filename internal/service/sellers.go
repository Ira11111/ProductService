package service

import "log/slog"

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
