package repository

import (
	"codebase-app/internal/module/shop/entity"
	"codebase-app/internal/module/shop/ports"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ShopRepository = &shopRepository{}

type shopRepository struct {
	db *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *shopRepository {
	return &shopRepository{
		db: db,
	}
}

func (r *shopRepository) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	var resp = new(entity.CreateShopResponse)
	// Your code here
	query := `
		INSERT INTO shops (user_id, name, description, terms)
		VALUES (?, ?, ?, ?) RETURNING id
	`

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.UserId,
		req.Name,
		req.Description,
		req.Terms).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateShop - Failed to create shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	var resp = new(entity.GetShopResponse)

	shopQuery := `
		SELECT name, description, terms
		FROM shops
		WHERE id = $1
	`

	err := r.db.QueryRowxContext(ctx, shopQuery, req.Id).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetShop - Failed to get shop")
		return nil, err
	}

	productQuery := `
		SELECT id, name, description, price, stock, image_url
		FROM products
		WHERE shop_id = $1
		AND deleted_at IS NULL
	`

	products := []entity.ProductItem{}
	err = r.db.SelectContext(ctx, &products, productQuery, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("repository::GetShop - Failed to get products")
		return nil, err
	}

	resp.Products = products

	return resp, nil
}

func (r *shopRepository) DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error {
	query := `
		UPDATE shops
		SET deleted_at = NOW()
		WHERE id = ? AND user_id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id, req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteShop - Failed to delete shop")
		return err
	}

	return nil
}

func (r *shopRepository) UpdateShop(ctx context.Context, req *entity.UpdateShopRequest) (*entity.UpdateShopResponse, error) {
	var resp = new(entity.UpdateShopResponse)

	query := `
		UPDATE shops
		SET name = ?, description = ?, terms = ?, updated_at = NOW()
		WHERE id = ? AND user_id = ?
		RETURNING id
	`

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Description,
		req.Terms,
		req.Id,
		req.UserId).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateShop - Failed to update shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) GetShops(ctx context.Context, req *entity.ShopsRequest) (*entity.ShopsResponse, error) {
	type dao struct {
		TotalData int `db:"total_data"`
		entity.ShopItem
		Category string `db:"category_name"`
	}

	var (
		resp = new(entity.ShopsResponse)
		data = make([]dao, 0, req.Paginate)
	)
	resp.Items = make([]entity.ShopItem, 0, req.Paginate)

	query := `
		SELECT
			COUNT(id) OVER() as total_data,
			id,
			name
		FROM shops
		WHERE
			deleted_at IS NULL
			AND user_id = ?
		LIMIT ? OFFSET ?
	`

	err := r.db.SelectContext(ctx, &data, r.db.Rebind(query),
		req.UserId,
		req.Paginate,
		req.Paginate*(req.Page-1),
	)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetShops - Failed to get shops")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	shopIds := make([]string, len(data))
	for i, d := range data {
		shopIds[i] = d.Id
	}

	var products []entity.ProductItem
	productQuery := `
		SELECT
			p.id,
			p.image_url,
			p.name,
			p.description,
			p.price,
			p.stock,
			p.user_id,
			p.shop_id,
			c.name as category
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE
			p.shop_id IN (?)
			AND p.deleted_at IS NULL
	`
	query, args, err := sqlx.In(productQuery, shopIds)
	if err != nil {
		log.Error().Err(err).Msg("repository::GetShops - Failed to construct product query")
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		log.Error().Err(err).Msg("repository::GetShops - Failed to get products")
		return nil, err
	}

	productMap := make(map[string][]entity.ProductItem)
	for _, product := range products {
		productMap[product.ShopId] = append(productMap[product.ShopId], product)
	}

	for _, d := range data {
		shop := d.ShopItem
		shop.Products = productMap[shop.Id]
		resp.Items = append(resp.Items, shop)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *shopRepository) CreateProduct(ctx context.Context, req *entity.CreateProductRequest) (*entity.CreateProductResponse, error) {
	var resp entity.CreateProductResponse
	query := `
        INSERT INTO products (shop_id, name, description, price, stock, user_id, image_url, category_id, merk, rating)
        VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?) 
        RETURNING id, shop_id, name, description, price, stock, user_id, category_id, image_url, merk, rating
    `

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.ShopId,
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.UserId,
		req.ImageURL,
		req.CategoryId,
		req.Merk,
		req.Rating,
	).Scan(&resp.Id, &resp.ShopId, &resp.Name, &resp.Description, &resp.Price, &resp.Stock, &resp.UserId, &resp.CategoryId, &resp.ImageURL, &resp.Merk, &resp.Rating)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateProduct - Failed to create product")
		return nil, err
	}

	return &resp, nil
}

func (r *shopRepository) GetProduct(ctx context.Context, req *entity.GetProductRequest) (*entity.GetProductResponse, error) {
	type dao struct {
		TotalData int    `db:"total_data"`
		Category  string `db:"category_name"`
		entity.ProductItem
	}

	var (
		resp = new(entity.GetProductResponse)
		data = make([]dao, 0, req.Paginate)
	)
	resp.ProductItem = make([]entity.ProductItem, 0, req.Paginate)

	query := `
		SELECT
			COUNT(p.id) OVER() as total_data,
			p.id,
			p.image_url,
			p.name,
			p.description,
			p.price,
			p.stock,
			p.user_id,
			p.shop_id,
			p.rating,
			p.merk,
			c.name as category_name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.deleted_at IS NULL
	`

	var args []interface{}
	if req.Keyword != "" {
		query += " AND p.name ILIKE ?"
		args = append(args, "%"+req.Keyword+"%")
	}
	if req.Category != "" {
		query += " AND c.name ILIKE ?"
		args = append(args, "%"+req.Category+"%")
	}
	if req.Rating != 0 {
		query += " AND p.rating = ?"
		args = append(args, req.Rating)
	}

	if req.Merk != "" {
		query += " AND p.merk ILIKE ?"
		args = append(args, "%"+req.Merk+"%")
	}

	query += `
		LIMIT ? OFFSET ?
	`

	args = append(args, req.Paginate, req.Paginate*(req.Page-1))
	err := r.db.SelectContext(ctx, &data, r.db.Rebind(query), args...)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProduct - Failed to get products")
		return nil, err
	}

	if len(data) > 0 {
		resp.Meta.TotalData = data[0].TotalData
	}

	for _, d := range data {
		product := d.ProductItem
		product.Category = d.Category
		resp.ProductItem = append(resp.ProductItem, product)
	}

	resp.Meta.CountTotalPage(req.Page, req.Paginate, resp.Meta.TotalData)

	return resp, nil
}

func (r *shopRepository) GetProductByid(ctx context.Context, req *entity.GetProductIdRequest) (*entity.GetProductIdResponse, error) {
	var resp = new(entity.GetProductIdResponse)
	query := `
		SELECT 
			id,
			name,
			price,
			stock,
			shop_id,
			user_id,
			image_url,
			description
		FROM products
		WHERE id = ?
		AND deleted_at IS NULL
	`

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query), req.Id).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetProduct - Failed to get product")
		return nil, err
	}
	return resp, nil
}

func (r *shopRepository) UpdateProduct(ctx context.Context, req *entity.UpdateProductRequest) (*entity.UpdateProductResponse, error) {
	var resp = new(entity.UpdateProductResponse)

	query := `
		UPDATE products
		SET name = ?, description = ?, price = ?, stock = ?, 
		    image_url = COALESCE(NULLIF(?, ''), image_url), 
		    updated_at = NOW()
		WHERE id = ? AND user_id = ?
		RETURNING id
	`
	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query),
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.ImageURL,
		req.Id,
		req.UserId).Scan(&resp.Id)

	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::UpdateProduct - Failed to update product")
		return nil, err
	}
	return resp, nil
}

func (r *shopRepository) DeleteProduct(ctx context.Context, req *entity.DeleteProductRequest) error {
	query := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = ? AND user_id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id, req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteProduct - Failed to delete product")
		return err
	}

	log.Info().Msg("repository::DeleteProduct - Product marked as deleted successfully")
	return nil
}

func (r *shopRepository) CreateCategory(ctx context.Context, req *entity.CreateCategoryRequest) (*entity.CreateCategoryResponse, error) {
	var resp = new(entity.CreateCategoryResponse)
	query := `
        INSERT INTO categories (name , user_id )
        VALUES (?, ?) RETURNING id, name, user_id
    `
	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query),
		req.Name,
		req.UserId,
	).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::CreateCategory - Failed to create category")
		return nil, err
	}
	return resp, nil
}

func (r *shopRepository) GetCategory(ctx context.Context, req *entity.GetCategoryRequest) (*entity.GetCategoryResponse, error) {
	var (
		items []entity.CategoryItem
		resp  = &entity.GetCategoryResponse{}
	)
	query := `
		SELECT id, name, user_id
		FROM categories
		WHERE user_id = ?
		AND deleted_at IS NULL
	`

	err := r.db.SelectContext(ctx, &items, r.db.Rebind(query), req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetCategory - Failed to get categories")
		return nil, err
	}

	resp.CategoryItems = items
	return resp, nil
}

func (r *shopRepository) GetCategoryId(ctx context.Context, req *entity.GetCategoryIdRequest) (*entity.GetcategoryIdResponse, error) {
	var resp entity.GetcategoryIdResponse

	query := `
        SELECT id, name, user_id
        FROM categories
        WHERE id = ?
		AND deleted_at IS NULL
    `

	err := r.db.QueryRowContext(ctx, query, req.Id).Scan(&resp.Id, &resp.Name, &resp.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::GetCategoryId - Failed to get category")
		return nil, err
	}

	return &resp, nil
}

func (r *shopRepository) DeletCategory(ctx context.Context, req *entity.DeleteCategoryRequest) error {

	query := `
		UPDATE categories
		SET deleted_at = NOW()
		WHERE id = ? AND user_id = ?
	`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id, req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("repository::DeleteCategory - Failed to delete category")
		return err
	}

	log.Info().Msg("repository::DeleteCategory - Category marked as deleted successfully")
	return nil
}
