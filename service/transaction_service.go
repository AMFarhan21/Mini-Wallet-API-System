package service

import (
	"context"
	"mini/model"

	"gorm.io/gorm"
)

type (
	TransactionRepo interface {
		Create(tx *gorm.DB, data model.Transactions) error
		GetHistory(ctx context.Context, userID, offset, limit int) ([]model.Transactions, error)
	}
)
