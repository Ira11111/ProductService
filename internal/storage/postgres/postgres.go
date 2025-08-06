package postgres

import (
	"github.com/Ira11111/ProductService/internal/config"
)

type Storage struct {
}

func NewStorage(dbCfg *config.DBConfig) (*Storage, error) {
	return &Storage{}, nil
}
