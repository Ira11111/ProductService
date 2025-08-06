package app

import (
	http "github.com/Ira11111/ProductService/internal/app/http"
	"github.com/Ira11111/ProductService/internal/config"
	server "github.com/Ira11111/ProductService/internal/http/products"
	"github.com/Ira11111/ProductService/internal/service"
	c "github.com/Ira11111/ProductService/internal/service/categories"
	p "github.com/Ira11111/ProductService/internal/service/products"
	s "github.com/Ira11111/ProductService/internal/service/sellers"
	w "github.com/Ira11111/ProductService/internal/service/warehouses"
	storage "github.com/Ira11111/ProductService/internal/storage/postgres"
	"log/slog"
	"time"
)

type App struct {
	logger         *slog.Logger
	HttpServer     *http.HTTPApp
	ProductService *service.ServiceAPI
}

func NewApp(logger *slog.Logger, dbCfg *config.DBConfig, port string, readTimeout time.Duration, writeTimeout time.Duration) *App {
	st, err := storage.NewStorage(dbCfg)
	if err != nil {
		panic(err)
	}

	productService := p.NewProductService(logger, st)
	sellerService := s.NewSellerService(logger, st)
	categoryService := c.NewCategoryService(logger, st)
	warehouseService := w.NewWarehouseService(logger, st)

	serviceApi := service.NewService(productService, sellerService, warehouseService, categoryService)

	serverApi := server.NewServerAPI(serviceApi)

	httpApp := http.NewHTTPApp(serverApi, port, readTimeout, writeTimeout)

	return &App{
		logger:         logger,
		HttpServer:     httpApp,
		ProductService: serviceApi,
	}
}
