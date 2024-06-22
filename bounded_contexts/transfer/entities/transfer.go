package entities

import (
	"errors"

	"github.com/shopspring/decimal"
)

type Transfer struct {
	Id                 int64
	SourceAccount      *Account
	DestinationAccount *Account
	Amount             decimal.Decimal
	Transaction        *Transaction
}

func NewTransfer(sourceAccount *Account, destinationAccount *Account, amount string) (*Transfer, error) {
	// Check source and dest account validity
	if sourceAccount == nil || destinationAccount == nil {
		return nil, errors.New("invalid source or destination account")
	}

	// Check if amount is valid
	amo, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, errors.New("invalid transfer amount")
	}
	if amo.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("transfer amount must be greater than zero")
	}

	// Check if source account has enough balance
	if sourceAccount.Balance.LessThan(amo) {
		return nil, errors.New("source account has insufficient balance")
	}

	// Update balance
	sourceAccount.Balance = sourceAccount.Balance.Sub(amo)
	destinationAccount.Balance = destinationAccount.Balance.Add(amo)

	// Generate receipt
	transaction := NewTransaction(sourceAccount.Id, destinationAccount.Id, amo)

	// Return
	return &Transfer{
		SourceAccount:      sourceAccount,
		DestinationAccount: destinationAccount,
		Amount:             amo,
		Transaction:        transaction,
	}, nil
}
