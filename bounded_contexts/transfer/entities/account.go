package entities

import "github.com/shopspring/decimal"

type Account struct {
	Id      int64
	Balance decimal.Decimal
}
