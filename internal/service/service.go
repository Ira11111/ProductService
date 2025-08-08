package service

import (
	"errors"
	storage "github.com/Ira11111/ProductService/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"log/slog"
	"slices"
)

var (
	adminRole               = "admin"
	sellerRole              = "seller"
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidRole          = errors.New("invalid role")
	ErrInvalidPermission    = errors.New("invalid permission")
	ErrInvalidUserId        = errors.New("invalid user id")
	ErrFailedToSaveEntity   = errors.New("failed to save entity")
	ErrFailedToDeleteEntity = errors.New("failed to delete entity")
	ErrFailedToUpdateEntity = errors.New("failed to update entity")
	ErrFailedToFindEntity   = errors.New("failed to find entity")
)

type ServiceAPI struct {
	logger  *slog.Logger
	storage *storage.Storage
}

func NewService(l *slog.Logger, s *storage.Storage) *ServiceAPI {
	return &ServiceAPI{
		logger:  l,
		storage: s,
	}
}

func (s *ServiceAPI) tokenInfo(c *gin.Context) (int64, []string, error) {
	rolesInterface, exists := c.Get("userRoles")
	if !exists {
		return 0, nil, ErrInvalidToken
	}

	roles, ok := rolesInterface.([]string)
	if !ok {
		return 0, nil, ErrInvalidRole
	}

	userInterface, exists := c.Get("userId")
	if !exists {
		return 0, nil, ErrInvalidToken
	}

	userId, ok := userInterface.(int64)
	if !ok {
		return 0, nil, ErrInvalidUserId
	}

	return userId, roles, nil
}

func (s *ServiceAPI) isAdmin(c *gin.Context) (bool, error) {
	const op = "service.IsAdmin"
	log := s.logger.With(slog.String("op", op))
	_, roles, err := s.tokenInfo(c)
	if err != nil {
		log.Warn("failed to get token")
		return false, err
	}
	if !slices.Contains(roles, adminRole) {
		log.Warn("permission denied")
		return false, ErrInvalidPermission
	}
	return true, nil
}
