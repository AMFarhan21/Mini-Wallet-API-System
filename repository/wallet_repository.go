package repository

import (
	"errors"
	"mini/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (r *WalletRepository) LockWallet(tx *gorm.DB, userID int) (*model.Wallets, error) {
	var wallet model.Wallets

	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("user_id=?", userID).First(&wallet).Error
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) Update(tx *gorm.DB, data model.Wallets) error {
	row := tx.Model(&model.Wallets{}).Where("id=?", data.ID).Update("balance", data.Balance)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("wallets not found")
	}

	return nil
}
