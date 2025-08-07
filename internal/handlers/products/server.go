package products

import (
	"github.com/Ira11111/ProductService/internal/service"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

var (
	adminRole  = "admin"
	sellerRole = "seller"
)

type ProductService interface {
	Products(c *gin.Context, offset int64, limit int64) (*api.ProductResponse, error)
	CreateProduct(c *gin.Context, product *api.ProductCreate, sellerId int64) (*api.ProductResponse, error)
	Product(c *gin.Context, id int64) (*api.ProductResponse, error)
	DeleteProduct(c *gin.Context, productId int64, sellerId int64) error
}

type SellerService interface{}

type WarehouseService interface{}

type CategoriesService interface {
	Categories(c *gin.Context) ([]*api.Category, error)
	CreateCategory(c *gin.Context) (*api.Category, error)
	DeleteCategory(c *gin.Context, id int) error
}

type ServerAPI struct {
	serviceApi *service.ServiceAPI
}

func NewServerAPI(ser *service.ServiceAPI) *ServerAPI {
	return &ServerAPI{
		serviceApi: ser,
	}
}

func (s *ServerAPI) GetCategories(c *gin.Context) {}

func (s *ServerAPI) PostCategories(c *gin.Context) {
}

func (s *ServerAPI) DeleteCategoriesId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) GetProducts(c *gin.Context, params api.GetProductsParams) {}

func (s *ServerAPI) PostProducts(c *gin.Context) {
	//TODO: обработка и сохранение фотографий
	rolesInterface, exists := c.Get("userRoles")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user roles not found"})
		return
	}
	roles, ok := rolesInterface.([]string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid roles type"})
		return
	}
	if !slices.Contains(roles, sellerRole) && !slices.Contains(roles, adminRole) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role for this action"})
		return
	}
	userInterface, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user id not found"})
		return
	}
	userId, ok := userInterface.(int64)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	var product api.ProductCreate
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respProducts, err := s.serviceApi.ProductProvider.CreateProduct(c, &product, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, respProducts)
}

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
