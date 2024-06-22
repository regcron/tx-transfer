package entities_test

import (
	"testing"

	"com/txfer/bounded_contexts/account/entities"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	initialBalance := "100.123"
	accountId := int64(1)

	account, err := entities.NewAccount(initialBalance, accountId)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if account.Id != accountId {
		t.Errorf("expected account Id to be %d, got %d", accountId, account.Id)
	}

	expectedBalance, _ := decimal.NewFromString(initialBalance)
	if account.Balance.Cmp(expectedBalance) != 0 {
		t.Errorf("expected account Balance to be %s, got %s", expectedBalance.String(), account.Balance.String())
	}
}
func TestInvalidAccountId(t *testing.T) {
	initialBalance := "100.123"
	accountId := int64(-1)

	account, err := entities.NewAccount(initialBalance, accountId)
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.EqualError(t, err, "invalid account id")
}

func TestInvalidInitialAmount(t *testing.T) {
	initialBalance := "abc"
	accountId := int64(1)

	account, err := entities.NewAccount(initialBalance, accountId)
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.EqualError(t, err, "invalid initial balance")

	initialBalance2 := "-100.231"
	account2, err2 := entities.NewAccount(initialBalance2, accountId)
	assert.NotNil(t, err2)
	assert.Nil(t, account2)
	assert.EqualError(t, err2, "initial balance must be greater than zero")
}

func TestParseAccountId(t *testing.T) {
	accountIdStr := "123"
	accountId, err := entities.ParseAccountId(accountIdStr)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), accountId)

	accountIdStr2 := "abc"
	accountId2, err2 := entities.ParseAccountId(accountIdStr2)
	assert.NotNil(t, err2)
	assert.Equal(t, int64(0), accountId2)
	assert.EqualError(t, err2, "invalid account id")

	accountIdStr3 := "-1"
	accountId3, err3 := entities.ParseAccountId(accountIdStr3)
	assert.NotNil(t, err3)
	assert.Equal(t, int64(0), accountId3)
	assert.EqualError(t, err3, "invalid account id")
}