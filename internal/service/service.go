package service

type ServiceAPI struct {
	ProductProvider    *ProductService
	SellerProvider     *SellerService
	WarehouseProvider  *WarehouseService
	CategoriesProvider *CategoryService
}

func NewService(productProvider *ProductService, sellerProvider *SellerService, warehouseProvider *WarehouseService, categoriesProvider *CategoryService) *ServiceAPI {
	return &ServiceAPI{
		ProductProvider:    productProvider,
		SellerProvider:     sellerProvider,
		WarehouseProvider:  warehouseProvider,
		CategoriesProvider: categoriesProvider,
	}
}
