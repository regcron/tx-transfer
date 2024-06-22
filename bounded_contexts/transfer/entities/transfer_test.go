package entities

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewTransfer(t *testing.T) {
	sourceAccount := &Account{Id: 1, Balance: decimal.NewFromFloat(100.123)}
	destinationAccount := &Account{Id: 2, Balance: decimal.NewFromFloat(0)}
	amount := "50.123"

	transfer, err := NewTransfer(sourceAccount, destinationAccount, amount)

	assert.NoError(t, err)
	assert.NotNil(t, transfer)
	assert.Equal(t, sourceAccount, transfer.SourceAccount)
	assert.Equal(t, destinationAccount, transfer.DestinationAccount)
	assert.Equal(t, decimal.NewFromFloat(50.123), transfer.Amount)
	assert.Equal(t, transfer.SourceAccount.Balance.Equal(decimal.NewFromFloat(50.0)), true)
	assert.Equal(t, transfer.DestinationAccount.Balance.Equal(decimal.NewFromFloat(50.123)), true)

	// Transaction receipt
	assert.NotNil(t, transfer.Transaction)
	assert.Equal(t, sourceAccount.Id, transfer.Transaction.SrcAccId)
	assert.Equal(t, destinationAccount.Id, transfer.Transaction.DestAccId)
	assert.Equal(t, transfer.Amount, transfer.Transaction.Amount)
	assert.NotNil(t, transfer.Transaction.Txid)
}

func TestNewTransfer_InvalidAmount(t *testing.T) {
	sourceAccount := &Account{Id: 1, Balance: decimal.NewFromFloat(100)}
	destinationAccount := &Account{Id: 2, Balance: decimal.NewFromFloat(0)}
	amount := "invalid"

	transfer, err := NewTransfer(nil, destinationAccount, amount)
	assert.Error(t, err)
	assert.Nil(t, transfer)
	assert.EqualError(t, err, "invalid source or destination account")

	transfer, err = NewTransfer(sourceAccount, destinationAccount, amount)

	assert.Error(t, err)
	assert.Nil(t, transfer)
	assert.EqualError(t, err, "invalid transfer amount")

	amount = "-1"
	transfer, err = NewTransfer(sourceAccount, destinationAccount, amount)
	assert.Error(t, err)
	assert.Nil(t, transfer)
	assert.EqualError(t, err, "transfer amount must be greater than zero")
}

func TestNewTransfer_InsufficientBalance(t *testing.T) {
	sourceAccount := &Account{Id: 1, Balance: decimal.NewFromFloat(10)}
	destinationAccount := &Account{Id: 2, Balance: decimal.NewFromFloat(0)}
	amount := "50.00"

	transfer, err := NewTransfer(sourceAccount, destinationAccount, amount)

	assert.Error(t, err)
	assert.Nil(t, transfer)
	assert.EqualError(t, err, "source account has insufficient balance")
}
