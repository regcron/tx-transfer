package dtos

import "com/txfer/bounded_contexts/account/entities"

type AccountResponse struct {
	Id      int64  `json:"account_id"`
	Balance string `json:"balance"`
}

func AccountResp(entity *entities.Account) *AccountResponse {
	return &AccountResponse{
		Id:      entity.Id,
		Balance: entity.Balance.String(),
	}
}
