package entity

import (
	"codebase-app/pkg/types"
)

type CreateShopRequest struct {
	UserId string `validate:"uuid" db:"user_id"`

	Name        string `json:"name" validate:"required" db:"name"`
	Description string `json:"description" validate:"required,max=255" db:"description"`
	Terms       string `json:"terms" validate:"required" db:"terms"`
}

type CreateShopResponse struct {
	Id string `json:"id" db:"id"`
}

type GetShopRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type GetShopResponse struct {
	Name        string        `json:"name" db:"name"`
	Description string        `json:"description" db:"description"`
	Terms       string        `json:"terms" db:"terms"`
	Products    []ProductItem `json:"products"`
}

type DeleteShopRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`

	Id string `validate:"uuid" db:"id"`
}

type UpdateShopRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`

	Id          string `params:"id" validate:"uuid" db:"id"`
	Name        string `json:"name" validate:"required" db:"name"`
	Description string `json:"description" validate:"required" db:"description"`
	Terms       string `json:"terms" validate:"required" db:"terms"`
}

type UpdateShopResponse struct {
	Id string `json:"id" db:"id"`
}

type ShopsRequest struct {
	UserId   string `prop:"user_id" validate:"uuid"`
	Page     int    `query:"page" validate:"required"`
	Paginate int    `query:"paginate" validate:"required"`
}

func (r *ShopsRequest) SetDefault() {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Paginate < 1 {
		r.Paginate = 10
	}
}

type ShopItem struct {
	Id       string        `json:"id" db:"id"`
	Name     string        `json:"name" db:"name"`
	Products []ProductItem `gorm:"foreignKey:ShopID"`
}

type ShopsResponse struct {
	Items []ShopItem `json:"items"`
	Meta  types.Meta `json:"meta"`
}

type CreateProductRequest struct {
	UserId      string `json:"userId" validate:"required"`
	ShopId      string `json:"shopId" validate:"required"`
	CategoryId  string `json:"categoryId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Merk        string `json:"merk" validate:"required"`
	Rating      int    `json:"rating" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
	ImageURL    string `json:"imageUrl"`
}

type CreateProductResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Rating      int     `json:"rating"`
	Merk        string  `json:"merk"`
	UserId      string  `validate:"uuid" db:"user_id"`
	ShopId      string  `validate:"uuid" db:"shop_id"`
	CategoryId  string  `validate:"uuid" db:"category_id"`
	ImageURL    string  `json:"imageUrl"`
}

type GetProductResponse struct {
	ProductItem []ProductItem
	Meta        types.Meta `json:"meta"`
}

type GetProductRequest struct {
	Page     int    `query:"page" validate:"required"`
	Paginate int    `query:"paginate" validate:"required"`
	Keyword  string `query:"keyword"`
	Rating   int    `query:"rating"`
	Merk     string `query:"merk"`
	Category string `query:"category"`
}

func (r *GetProductRequest) SetDefault() {
	if r.Page < 1 {
		r.Page = 1
	}

	if r.Paginate < 1 {
		r.Paginate = 10
	}
}

type ProductItem struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	ShopId      string  `json:"shopId" db:"shop_id"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	Merk        string  `json:"merk" db:"merk"`
	Rating      int     `json:"rating" db:"rating"`
	Stock       int     `json:"stock" db:"stock"`
	UserId      string  `json:"userId" db:"user_id"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
}

type GetProductIdRequest struct {
	Id string `validate:"uuid" db:"id"`
}

type GetProductIdResponse struct {
	Id          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	ShopId      string  `json:"shopId" db:"shop_id"`
	UserId      string  `json:"userId" db:"user_id"`
	Price       float64 `json:"price" db:"price"`
	Stock       int     `json:"stock" db:"stock"`
	Description string  `json:"description" db:"description"`
	ImageURL    string  `json:"imageUrl" db:"image_url"`
}

type UpdateProductRequest struct {
	UserId      string `prop:"user_id" validate:"uuid" db:"user_id"`
	Id          string `params:"id" validate:"uuid" db:"id"`
	Name        string `json:"name" validate:"required" db:"name"`
	Description string `json:"description" validate:"required" db:"description"`
	Price       int    `json:"price" validate:"required" db:"price"`
	Stock       int    `json:"stock" validate:"required" db:"stock"`
	ImageURL    string `json:"imageUrl" db:"image_url"`
}

type UpdateProductResponse struct {
	Id string `json:"id" db:"id"`
}

type DeleteProductRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`
	Id     string `validate:"uuid" db:"id"`
}

type CreateCategoryRequest struct {
	UserId string `json:"userId" validate:"required"`
	Name   string `json:"name" validate:"required" db:"name"`
}

type CreateCategoryResponse struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserId string `json:"userId" db:"user_id"`
}

type CategoryItem struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserId string `json:"userId" db:"user_id"`
}

type GetCategoryResponse struct {
	CategoryItems []CategoryItem `json:"categoryItems"`
}

type GetCategoryRequest struct {
	UserId string `json:"user_id" validate:"required,uuid"`
}

type GetCategoryIdRequest struct {
	Id     string `json:"id" validate:"required,uuid"`
	UserId string `json:"user_id" validate:"required,uuid"`
}

type GetcategoryIdResponse struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserId string `json:"userId" db:"user_id"`
}

type DeleteCategoryRequest struct {
	UserId string `prop:"user_id" validate:"uuid" db:"user_id"`
	Id     string `validate:"uuid" db:"id"`
}
