package http

import (
	"context"
	"errors"
	"github.com/Ira11111/ProductService/internal/handlers/products"
	m "github.com/Ira11111/go-interceptors/midlewares/gin"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type HTTPApp struct {
	ginEngine  *gin.Engine
	httpServer *http.Server
	serverApi  *products.ServerAPI
	port       string
}

func NewHTTPApp(server *products.ServerAPI, port string, readTimeout time.Duration, writeTimeout time.Duration) *HTTPApp {
	engine := gin.Default()
	engine.Use(gin.Logger())
	key := os.Getenv("JWT_PUBLIC_KEY")
	if key == "" {
		panic("jwt public key not found")
	}
	midlleware := m.NewAuthMiddleware(key)
	api.RegisterHandlersWithOptions(
		engine,
		server,
		api.GinServerOptions{
			Middlewares: []api.MiddlewareFunc{
				api.MiddlewareFunc(midlleware.JWTClaims()),
			},
		},
	)

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return &HTTPApp{
		ginEngine:  engine,
		httpServer: httpServer,
		serverApi:  server,
		port:       port,
	}
}

func (app *HTTPApp) Start() error {
	if err := app.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (app *HTTPApp) Stop(ctx context.Context) {
	if err := app.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
