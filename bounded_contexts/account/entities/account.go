package entities

import (
	"errors"
	"strconv"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id      int64
	Balance decimal.Decimal
}

func NewAccount(initialBalance string, accountId int64) (*Account, error) {
	if accountId <= 0 {
		return nil, errors.New("invalid account id")
	}

	bal, err := decimal.NewFromString(initialBalance)
	if err != nil {
		return nil, errors.New("invalid initial balance")
	}

	if bal.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("initial balance must be greater than zero")
	}

	return &Account{
		Id:      accountId,
		Balance: bal,
	}, nil
}

func ParseAccountId(accountIdStr string) (int64, error) {
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		return 0, errors.New("invalid account id")
	}
	if accountId <= 0 {
		return 0, errors.New("invalid account id")
	}
	return int64(accountId), nil
}
