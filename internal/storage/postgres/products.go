package postgres

import (
	"errors"
	"fmt"
	"github.com/Ira11111/ProductService/internal/storage"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (s *Storage) Products(c *gin.Context, offset int64, limit int64) ([]*api.ProductResponse, error) {
	return nil, nil
}
func (s *Storage) SaveProduct(c *gin.Context, product *api.ProductCreate, sellerId int64) (*api.ProductResponse, error) {
	const op = "storage.postgres.SaveProduct"
	tx, err := s.db.BeginTx(c, nil)
	if err != nil {
		return nil, storage.ErrFailedStartTransaction
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(
		c,
		"INSERT INTO products (name, description, price, seller_id) VALUES ($1, $2, $3, $4) RETURNING *",
		product.Name, product.Description, product.Price, sellerId,
	)

	var newProduct api.ProductResponse
	err = row.Scan(&newProduct.Id, &newProduct.Name, &newProduct.Description, &newProduct.Price, &newProduct.Seller)
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

	stmt, err := tx.Prepare("INSERT INTO category_product product_id, category_id VALUES ($1, $2)")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
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
	return nil, nil
}
func (s *Storage) DeleteProduct(c *gin.Context, id int64) error {
	return nil
}
