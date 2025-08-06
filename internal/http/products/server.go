package products

import (
	"github.com/Ira11111/ProductService/internal/service"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

type ServerAPI struct {
	serviceApi *service.ServiceAPI
}

func NewServerAPI(ser *service.ServiceAPI) *ServerAPI {
	return &ServerAPI{
		serviceApi: ser,
	}
}

func (s *ServerAPI) GetCategories(c *gin.Context) {}

func (s *ServerAPI) GetProducts(c *gin.Context, params api.GetProductsParams) {}

func (s *ServerAPI) PostProducts(c *gin.Context) {}

func (s *ServerAPI) DeleteProductsId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) GetProductsId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PutProductsId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PostSellers(c *gin.Context) {}

func (s *ServerAPI) DeleteSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) GetSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PutSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PostWarehouse(c *gin.Context) {}

func (s *ServerAPI) PostWarehouseIdProducts(c *gin.Context, id api.IdParam, params api.PostWarehouseIdProductsParams) {
}

func (s *ServerAPI) GetWarehousesId(c *gin.Context, id api.IdParam) {}
