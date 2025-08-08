package postgres

import (
	"database/sql"
	"errors"
	"github.com/Ira11111/ProductService/internal/storage"
	api "github.com/Ira11111/protos/v4/gen/go/products"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (s *Storage) Categories(c *gin.Context) ([]*api.Category, error) {
	rows, err := s.db.Query(`SELECT id, name FROM categories`)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrEntityNotFound
		}
		return nil, err
	}
	res := make([]*api.Category, 0)

	for rows.Next() {
		var curCategory api.Category
		if err = rows.Scan(&curCategory.Id, &curCategory.Name); err != nil {
			return nil, err
		}
		res = append(res, &curCategory)
	}
	return res, nil
}

func (s *Storage) SaveCategory(c *gin.Context, categoryName string) (*api.Category, error) {
	row := s.db.QueryRowContext(c, `INSERT INTO categories (name) VALUES ($1) RETURNING id, name`, categoryName)
	var category api.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == ErrInvalidFormat {
				return nil, storage.ErrFailedToSaveEntity
			}
		}
		return nil, err
	}
	return &category, nil
}

func (s *Storage) DropCategory(c *gin.Context, id int64) error {
	_, err := s.db.ExecContext(c, `DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.ErrEntityNotFound
		}
		return err
	}
	return nil
}
