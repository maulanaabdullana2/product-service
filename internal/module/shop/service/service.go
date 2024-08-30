package service

import (
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	"context"
)

var _ ports.ShopService = &shopService{}

type shopService struct {
	repo ports.ShopRepository
}

func NewShopService(repo ports.ShopRepository) *shopService {
	return &shopService{
		repo: repo,
	}
}

func (s *shopService) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	return s.repo.CreateShop(ctx, req)
}

func (s *shopService) GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	return s.repo.GetShop(ctx, req)
}

func (s *shopService) DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error {
	return s.repo.DeleteShop(ctx, req)
}

func (s *shopService) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
	return s.repo.UpdateShop(ctx, req)
}

func (s *shopService) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
	return s.repo.GetShops(ctx, req)
}

func (s *shopService) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	return s.repo.CreateProduct(ctx, req)
}

func (s *shopService) GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error) {
	return s.repo.GetProduct(ctx, req)
}

func (s *shopService) GetProdctByid(ctx context.Context, req *entity.GetProductIdRequest) (*entity.GetProductIdResponse, error) {
	return s.repo.GetProductByid(ctx, req)
}

func (s *shopService) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	return s.repo.UpdateProduct(ctx, req)
}

func (s *shopService) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	return s.repo.DeleteProduct(ctx, req)
}

func (s *shopService) CreateCategory(ctx context.Context, req *entity.CreateCategoryRequest) (*entity.CreateCategoryResponse, error) {
	return s.repo.CreateCategory(ctx, req)
}

func (s *shopService) GetCategory(ctx context.Context, req *entity.GetCategoryRequest) (*entity.GetCategoryResponse, error) {
	return s.repo.GetCategory(ctx, req)
}

func (s *shopService) GetCategoryId(ctx context.Context, req *entity.GetCategoryIdRequest) (*entity.GetcategoryIdResponse, error) {
	return s.repo.GetCategoryId(ctx, req)
}

func (s *shopService) DeletCategory(ctx context.Context, req *entity.DeleteCategoryRequest) error {
	return s.repo.DeletCategory(ctx, req)
}
