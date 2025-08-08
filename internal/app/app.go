package app

import (
	h "github.com/Ira11111/ProductService/internal/app/http"
	"github.com/Ira11111/ProductService/internal/config"
	server "github.com/Ira11111/ProductService/internal/handlers/products"
	s "github.com/Ira11111/ProductService/internal/service"
	storage "github.com/Ira11111/ProductService/internal/storage/postgres"
	"log/slog"
	"time"
)

type App struct {
	logger         *slog.Logger
	HttpServer     *h.HTTPApp
	ProductService *s.ServiceAPI
}

func NewApp(logger *slog.Logger, dbCfg *config.DBConfig, port string, readTimeout time.Duration, writeTimeout time.Duration) *App {
	st, err := storage.NewStorage(dbCfg)
	if err != nil {
		panic(err)
	}

	serviceApi := s.NewService(logger, st)
	serverApi := server.NewServerAPI(serviceApi)
	httpApp := h.NewHTTPApp(serverApi, port, readTimeout, writeTimeout)

	return &App{
		logger:         logger,
		HttpServer:     httpApp,
		ProductService: serviceApi,
	}
}
