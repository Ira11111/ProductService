package app

import (
	http "github.com/Ira11111/ProductService/internal/app/http"
	server "github.com/Ira11111/ProductService/internal/http/products"
	c "github.com/Ira11111/ProductService/internal/service/categories"
	p "github.com/Ira11111/ProductService/internal/service/products"
	s "github.com/Ira11111/ProductService/internal/service/sellers"
	w "github.com/Ira11111/ProductService/internal/service/warehouses"
	"log/slog"
)

type App struct {
	logger     *slog.Logger
	httpServer *http.HTTPApp
	port       int
}

func NewApp(logger *slog.Logger, port int) *App {
	productService := p.NewProductService(logger)
	sellerService := s.NewSellerService(logger)
	categoryService := c.NewCategoryService(logger)
	warehouseService := w.NewWarehouseService(logger)

	serverApi := server.NewServerAPI(productService, sellerService, categoryService, warehouseService)
	httpApp := http.NewHTTPApp(serverApi)

	return &App{
		logger:     logger,
		httpServer: httpApp,
		port:       port,
	}
}
