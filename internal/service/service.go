package service

type ServiceAPI struct {
	productProvider    ProductService
	sellerProvider     SellerService
	warehouseProvider  WarehouseService
	categoriesProvider CategoriesService
}

type ProductService interface {
	GetProduct() error
}

type SellerService interface{}

type WarehouseService interface{}

type CategoriesService interface{}

func NewService(productProvider ProductService, sellerProvider SellerService, warehouseProvider WarehouseService, categoriesProvider CategoriesService) *ServiceAPI {
	return &ServiceAPI{
		productProvider:    productProvider,
		sellerProvider:     sellerProvider,
		warehouseProvider:  warehouseProvider,
		categoriesProvider: categoriesProvider,
	}
}
