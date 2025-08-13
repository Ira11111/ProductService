package storage

import "errors"

var (
	ErrSellerNotFound         = errors.New("seller not found")
	ErrCategoryNotFound       = errors.New("category not found")
	ErrEntityNotFound         = errors.New("entity not found")
	ErrFailedStartTransaction = errors.New("failed to start transaction")
	ErrFailedToSaveEntity     = errors.New("failed to save entity")
	ErrCommitFailed           = errors.New("failed to commit")
	ErrInvalidQuery           = errors.New("invalid DB query")
)
