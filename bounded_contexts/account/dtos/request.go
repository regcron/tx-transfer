package dtos

type CreateAccountRequest struct {
	AccountId      int64  `json:"account_id" binding:"required"`
	InitialBalance string `json:"initial_balance" binding:"required"`
}
