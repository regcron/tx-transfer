package dtos

import (
	"com/txfer/bounded_contexts/transfer/entities"
	"time"

	"github.com/shopspring/decimal"
)

type TransactionDto struct {
	Id        int64           `gorm:"column:id"`
	Txid      string          `gorm:"column:tx_id"`
	SrcAccId  int64           `gorm:"column:src_acc_id"`
	DestAccId int64           `gorm:"column:dest_acc_id"`
	Amount    decimal.Decimal `gorm:"column:amount"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
}

func NewTransactionDto(txnEntity *entities.Transaction) *TransactionDto {
	return &TransactionDto{
		Txid:      txnEntity.Txid,
		SrcAccId:  txnEntity.SrcAccId,
		DestAccId: txnEntity.DestAccId,
		Amount:    txnEntity.Amount,
	}
}

func (dto *TransactionDto) ToEntity() *entities.Transaction {
	return &entities.Transaction{
		Id:        dto.Id,
		Txid:      dto.Txid,
		SrcAccId:  dto.SrcAccId,
		DestAccId: dto.DestAccId,
		Amount:    dto.Amount,
	}
}
func (TransactionDto) TableName() string {
	return "transactions"
}
