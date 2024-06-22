package entities

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Id        int64
	Txid      string
	SrcAccId  int64
	DestAccId int64
	Amount    decimal.Decimal
}

func NewTransaction(srcAccId int64, destAccId int64, amount decimal.Decimal) *Transaction {
	return &Transaction{
		Txid:      uuid.New().String(),
		SrcAccId:  srcAccId,
		DestAccId: destAccId,
		Amount:    amount,
	}
}
