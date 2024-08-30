package ports

import (
	"codebase-app/internal/module/shop/entity"
	"context"
)

type ShopRepository interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
	UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error)
	GetProductByid(ctx context.Context, req *entity.GetProductIdRequest) (*entity.GetProductIdResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
	CreateCategory(ctx context.Context, req *entity.CreateCategoryRequest) (*entity.CreateCategoryResponse, error)
	GetCategory(ctx context.Context, req *entity.GetCategoryRequest) (*entity.GetCategoryResponse, error)
	GetCategoryId(ctx context.Context, req *entity.GetCategoryIdRequest) (*entity.GetcategoryIdResponse, error)
	DeletCategory(ctx context.Context, req *entity.DeleteCategoryRequest) error
}

type ShopService interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
	UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error)
	GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error)
	CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error)
	GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error)
	GetProdctByid(ctx context.Context, req *entity.GetProductIdRequest) (*entity.GetProductIdResponse, error)
	UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error
	CreateCategory(ctx context.Context, req *entity.CreateCategoryRequest) (*entity.CreateCategoryResponse, error)
	GetCategory(ctx context.Context, req *entity.GetCategoryRequest) (*entity.GetCategoryResponse, error)
	GetCategoryId(ctx context.Context, req *entity.GetCategoryIdRequest) (*entity.GetcategoryIdResponse, error)
	DeletCategory(ctx context.Context, req *entity.DeleteCategoryRequest) error
}
