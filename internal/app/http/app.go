package http

import (
	"github.com/Ira11111/ProductService/internal/http/products"
	"github.com/gin-gonic/gin"
)

type HTTPApp struct {
	ginEngine *gin.Engine
	serverApi *products.ServerAPI
}

func NewHTTPApp(server *products.ServerAPI) *HTTPApp {
	engine := gin.New()
	return &HTTPApp{
		ginEngine: engine,
		serverApi: server,
	}
}
