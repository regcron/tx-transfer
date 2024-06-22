package repositories

import (
	"com/txfer/bounded_contexts/transfer/dtos"
	"com/txfer/bounded_contexts/transfer/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransferRepo struct {
	db *gorm.DB
}

func NewTransferRepo(db *gorm.DB) *TransferRepo {
	return &TransferRepo{db: db}
}

func (repo *TransferRepo) GetAccountByIdForUpdate(tx *gorm.DB, id int64) (*entities.Account, error) {
	var dto dtos.AccountDto
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&dto).Error; err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}

func (repo *TransferRepo) CreateTransfer(srcAccId int64, descAccId int64, amount string) (*entities.Transaction, error) {
	var txnDto *dtos.TransactionDto
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		// Create new transfer
		srcAcc, err := repo.GetAccountByIdForUpdate(tx, srcAccId)
		if err != nil {
			return err
		}
		destAcc, err := repo.GetAccountByIdForUpdate(tx, descAccId)
		if err != nil {
			return err
		}
		newTransfer, err := entities.NewTransfer(srcAcc, destAcc, amount)
		if err != nil {
			return err
		}

		// Update DB
		srcAccDto := dtos.NewAccountDto(srcAcc)
		destAccDto := dtos.NewAccountDto(destAcc)
		txnDto = dtos.NewTransactionDto(newTransfer.Transaction)

		if err := tx.Model(srcAccDto).Update("balance", srcAccDto.Balance).Error; err != nil {
			return err
		}
		if err := tx.Model(destAccDto).Update("balance", destAccDto.Balance).Error; err != nil {
			return err
		}
		if err := tx.Create(txnDto).Error; err != nil {
			return err
		}
		if err := tx.Model(txnDto).Where("tx_id = ?", txnDto.Txid).First(txnDto).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return txnDto.ToEntity(), nil
}
