package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Ira11111/ProductService/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib" // Вот правильный импорт для database/sql
	"github.com/lib/pq"
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
	db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		dbCfg.User,
		dbCfg.Pass,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Name,
	))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}
