package dtos

import "com/txfer/bounded_contexts/transfer/entities"

type TransferResponseDto struct {
	Id        string `json:"transaction_id"`
	SrcAccId  int64  `json:"source_account_id"`
	DestAccId int64  `json:"destination_account_id"`
	Amount    string `json:"amount"`
}

func TransferResp(txn *entities.Transaction) *TransferResponseDto {
	return &TransferResponseDto{
		Id:        txn.Txid,
		SrcAccId:  txn.SrcAccId,
		DestAccId: txn.DestAccId,
		Amount:    txn.Amount.String(),
	}
}
