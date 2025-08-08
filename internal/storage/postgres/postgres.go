package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Ira11111/ProductService/internal/config"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	ErrUniqueViolation     pq.ErrorCode = "23505"
	ErrForeignKeyViolation pq.ErrorCode = "23503"
	ErrInvalidFormat       pq.ErrorCode = "22P02"
	//ErrNotNullViolation    pq.ErrorCode = "23502"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(dbCfg *config.DBConfig) (*Storage, error) {
	const op = "storage.NewStorage"

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.User,
			dbCfg.Pass,
			dbCfg.Name,
			dbCfg.SSL,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
