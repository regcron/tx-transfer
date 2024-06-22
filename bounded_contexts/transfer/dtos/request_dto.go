package dtos

type TransferRequestDto struct {
	SrcAccId  int64  `json:"source_account_id" binding:"required"`
	DestAccId int64  `json:"destination_account_id" binding:"required"`
	Amount    string `json:"amount" binding:"required"`
}
