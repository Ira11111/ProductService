package products

import (
	"errors"
	"net/http"

	"github.com/Ira11111/ProductService/internal/service"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
)

const (
	defaultLimit  = 10
	defaultOffset = 0
)

type ProductService interface {
	CreateProduct(c *gin.Context, product *api.ProductCreate) (*api.ProductResponse, error)
	Product(c *gin.Context, id int64) (*api.ProductResponse, error)
	DeleteProduct(c *gin.Context, productId int64) error
	EditProduct(c *gin.Context, create *api.ProductCreate) (*api.ProductResponse, error)
}

type ProductListService interface {
	Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error)
	ProductsCategory(c *gin.Context, id int64, offset int64, limit int64) (*[]api.ProductResponse, error)
	ProductsWarehouse(c gin.Context, id int64, offset int64, limit int64) (*api.ProductResponse, error)
}

type SellerService interface {
	CreateSeller(c *gin.Context, seller *api.SellerFull) (*api.SellerFull, error)
	Seller(c *gin.Context, id int64) (*api.SellerFull, error)
	DeleteSeller(c *gin.Context, id int64) error
	EditSeller(c *gin.Context, full *api.SellerFull) (*api.SellerFull, error)
}

type WarehouseService interface {
	Warehouse(c *gin.Context, id int64) (*api.Warehouse, error)
	CreateWarehouse(c *gin.Context, warehouse *api.Warehouse) (*api.Warehouse, error)
}

type CategoriesService interface {
	Categories(c *gin.Context) ([]*api.Category, error)
	CreateCategory(c *gin.Context, category *api.Category) (*api.Category, error)
	DeleteCategory(c *gin.Context, id int64) error
}

type ServerAPI struct {
	serviceApi *service.ServiceAPI
}

func NewServerAPI(ser *service.ServiceAPI) *ServerAPI {
	return &ServerAPI{
		serviceApi: ser,
	}
}

func (s *ServerAPI) GetCategories(c *gin.Context) {
	categories, err := s.serviceApi.Categories(c)
	if err != nil {
		if errors.Is(err, service.ErrFailedToFindEntity) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no entities"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (s *ServerAPI) PostCategories(c *gin.Context) {
	var category api.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCat, err := s.serviceApi.CreateCategory(c, &category)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category": newCat})
}

func (s *ServerAPI) DeleteCategoriesId(c *gin.Context, id api.IdParam) {
	err := s.serviceApi.DeleteCategory(c, id)
	if err != nil {
		if errors.Is(err, service.ErrFailedToFindEntity) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "category not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "category deleted"})
}

func (s *ServerAPI) GetCategoriesIdProducts(c *gin.Context, id api.IdParam, params api.GetCategoriesIdProductsParams) {
	return
}

func (s *ServerAPI) GetProducts(c *gin.Context, params api.GetProductsParams) {
	var lim int64
	var off int64
	if params.Limit == nil {
		lim = defaultLimit
	} else {
		lim = *params.Limit
	}
	if params.Offset == nil {
		off = defaultOffset
	} else {
		off = *params.Offset
	}
	products, err := s.serviceApi.Products(c, off, lim)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"offset": off, "amount": lim, "products": products})
}

func (s *ServerAPI) PostProducts(c *gin.Context) {
	//TODO: обработка и сохранение фотографий

	var product api.ProductCreate
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respProducts, err := s.serviceApi.CreateProduct(c, &product)
	if err != nil {
		switch err {
		case service.ErrInvalidToken:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		case service.ErrInvalidUserId:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
			return
		case service.ErrInvalidPermission:
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid permission"})
			return
		case service.ErrInvalidRole:
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid role"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusCreated, respProducts)
	return
}

func (s *ServerAPI) DeleteProductsId(c *gin.Context, id api.IdParam) {
	err := s.serviceApi.DeleteProduct(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "product deleted successful"})
	return
}

func (s *ServerAPI) GetProductsId(c *gin.Context, id api.IdParam) {
	product, err := s.serviceApi.Product(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (s *ServerAPI) PutProductsId(c *gin.Context, id api.IdParam) {
	var editedProduct api.ProductCreate
	if err := c.ShouldBindJSON(&editedProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	product, err := s.serviceApi.EditProduct(c, &editedProduct)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})

}

func (s *ServerAPI) PostSellers(c *gin.Context) {}

func (s *ServerAPI) DeleteSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) GetSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PutSellersId(c *gin.Context, id api.IdParam) {}

func (s *ServerAPI) PostWarehouse(c *gin.Context) {}

func (s *ServerAPI) GetWarehouseIdProducts(c *gin.Context, id api.IdParam, params api.GetWarehouseIdProductsParams) {

}

func (s *ServerAPI) GetWarehousesId(c *gin.Context, id api.IdParam) {}
