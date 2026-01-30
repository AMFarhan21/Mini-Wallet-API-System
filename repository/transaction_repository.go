package repository

import (
	"context"
	"mini/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Create(tx *gorm.DB, data model.Transactions) error {
	return tx.Create(&data).Error
}

func (r *TransactionRepository) GetHistory(ctx context.Context, userID, offset, limit int) ([]model.Transactions, error) {
	var transaction []model.Transactions

	err := r.db.WithContext(ctx).
		Joins("JOIN wallets w ON w.id = transactions.wallet_id").
		Where("w.user_id=?", userID).
		Order("transactions.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
