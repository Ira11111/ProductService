package postgres

import api "github.com/Ira11111/protos/v4/gen/go/products"

func (s *Storage) Categories() ([]*api.Category, error) {
	return nil, nil
}

func (s *Storage) SaveCategory(categoryName string) (*api.Category, error) {
	return nil, nil
}

func (s *Storage) DropCategory(id int64) error {
	return nil
}
