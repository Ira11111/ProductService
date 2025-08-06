package categories

import "log/slog"

type CategoryProvider interface{}

type CategoryService struct {
	logger   *slog.Logger
	provider CategoryProvider
}

func NewCategoryService(logger *slog.Logger, provider CategoryProvider) *CategoryService {
	return &CategoryService{
		logger:   logger,
		provider: provider,
	}
}
