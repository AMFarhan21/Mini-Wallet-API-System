package service

import (
	"context"
	"errors"
	"mini/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	WalletRepo interface {
		LockWallet(tx *gorm.DB, userID int) (*model.Wallets, error)
		Update(tx *gorm.DB, data model.Wallets) error
	}

	WalletService struct {
		db              *gorm.DB
		walletRepo      WalletRepo
		transactionRepo TransactionRepo
	}
)

func NewWalletService(db *gorm.DB, walletRepo WalletRepo, transactionRepo TransactionRepo) *WalletService {
	return &WalletService{
		db:              db,
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *WalletService) TopUp(ctx context.Context, userID int, amount int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		wallet, err := s.walletRepo.LockWallet(tx, userID)
		if err != nil {
			return errors.New("wallet doesn't exists")
		}

		wallet.Balance += amount
		transaction := &model.Transactions{
			WalletId:    wallet.ID,
			Type:        "Credit",
			Amount:      amount,
			ReferenceId: uuid.NewString(),
		}

		err = s.walletRepo.Update(tx, *wallet)
		if err != nil {
			return err
		}

		err = s.transactionRepo.Create(tx, *transaction)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *WalletService) Transfer(ctx context.Context, fromUserID, toUserID int, amount int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if amount <= 0 {
			return errors.New("invalid amount")
		}

		if fromUserID == toUserID {
			return errors.New("invalid target id")
		}

		senderWallet, err := s.walletRepo.LockWallet(tx, fromUserID)
		if err != nil {
			return errors.New("user not found")
		}

		receiverWallet, err := s.walletRepo.LockWallet(tx, toUserID)
		if err != nil {
			return errors.New("user not found")
		}

		if senderWallet.Balance < amount {
			return errors.New("insufficient balance")
		}

		senderWallet.Balance -= amount
		receiverWallet.Balance += amount

		senderTransaction := &model.Transactions{
			WalletId:    senderWallet.ID,
			Type:        "Debit",
			ReferenceId: uuid.NewString(),
			Amount:      amount,
		}
		receiverTransaction := &model.Transactions{
			WalletId:    receiverWallet.ID,
			Type:        "Credit",
			ReferenceId: uuid.NewString(),
			Amount:      amount,
		}

		err = s.walletRepo.Update(tx, *senderWallet)
		if err != nil {
			return err
		}
		err = s.walletRepo.Update(tx, *receiverWallet)
		if err != nil {
			return err
		}

		err = s.transactionRepo.Create(tx, *senderTransaction)
		if err != nil {
			return err
		}
		err = s.transactionRepo.Create(tx, *receiverTransaction)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *WalletService) GetHistory(ctx context.Context, userID, page, limit int) ([]model.Transactions, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 5
	}

	offset := (page - 1) * limit

	return s.transactionRepo.GetHistory(ctx, userID, offset, limit)
}
