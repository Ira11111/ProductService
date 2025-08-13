package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Ira11111/ProductService/internal/storage"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

//TODO: вынести в отельную функцию создание списка продуктов

func (s *Storage) Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error) {
	res := make([]*api.ProductResponse, 0, limit)
	fmt.Println(limit, offset)
	rows, err := s.db.QueryContext(c, `SELECT p.id, p.name, p.description, p.price, s.id, s.name
    FROM products p INNER JOIN sellers s ON s.id = p.seller_id 
    ORDER BY p.order_count OFFSET $1 LIMIT $2`, offset, limit)

	if err != nil {
		return nil, storage.ErrInvalidQuery
	}

	for rows.Next() {
		var seller api.SellerShort
		var product api.ProductResponse
		photos := make([]string, 0)
		product.Seller = &seller
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &seller.Id, &seller.Name); err != nil {
			return nil, err
		}

		photoRow, err := s.db.Query(`SELECT file_path from product_images WHERE product_id = $1`, product.Id)
		if err != nil {
			return nil, err
		}
		for photoRow.Next() {
			var path string
			if err = photoRow.Scan(&path); err != nil {
				return nil, err
			}
			photos = append(photos, path)
		}
		product.PhotoUrls = &photos

		res = append(res, &product)
	}
	return res, nil
}

func (s *Storage) SaveProduct(c *gin.Context, product *api.ProductCreate, userId int64) (*api.ProductResponse, error) {
	const op = "storage.postgres.SaveProduct"
	tx, err := s.db.BeginTx(c, nil)
	if err != nil {
		return nil, storage.ErrFailedStartTransaction
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(c, "SELECT id, name from sellers WHERE user_id=$1", userId)
	var seller api.SellerShort
	if err = row.Scan(&seller.Id, &seller.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrSellerNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row = tx.QueryRowContext(
		c,
		"INSERT INTO products (name, description, price, seller_id) VALUES ($1, $2, $3, $4) RETURNING id, name, description, price",
		product.Name, product.Description, product.Price, seller.Id,
	)

	var newProduct api.ProductResponse
	err = row.Scan(&newProduct.Id, &newProduct.Name, &newProduct.Description, &newProduct.Price)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == ErrForeignKeyViolation {
				return nil, storage.ErrSellerNotFound
			}
			if pqErr.Code == ErrInvalidFormat {
				return nil, storage.ErrFailedToSaveEntity
			}
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := tx.Prepare("INSERT INTO category_product (product_id, category_id) VALUES ($1, $2)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	for _, catId := range *product.CategoriesId {
		_, err = stmt.ExecContext(c, newProduct.Id, catId)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) {
				if pqErr.Code == ErrForeignKeyViolation {
					return nil, storage.ErrCategoryNotFound
				}
			}
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	newProduct.CategoriesId = product.CategoriesId
	err = tx.Commit()
	if err != nil {
		return nil, storage.ErrCommitFailed
	}
	return &newProduct, nil
}
func (s *Storage) Product(c *gin.Context, id int64) (*api.ProductResponse, error) {
	//тут просто находим
	return nil, nil
}
func (s *Storage) DeleteProduct(c *gin.Context, id int64) error {
	// тут просто удаляем
	return nil
}

func (s *Storage) UpdateProduct(c *gin.Context, create *api.ProductCreate) (*api.ProductCreate, error) {
	// тут просто обновляем
	return nil, nil
}

func (s *Storage) ProductsByCategory(c *gin.Context, id int64, offset int64, limit int64) (*[]api.ProductResponse, error) {
	return nil, nil
}
func (s *Storage) ProductsByWarehouse(c gin.Context, id int64, offset int64, limit int64) (*api.ProductResponse, error) {
	return nil, nil
}
