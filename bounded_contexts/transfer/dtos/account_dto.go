package dtos

import (
	"com/txfer/bounded_contexts/transfer/entities"
	"time"

	"github.com/shopspring/decimal"
)

type AccountDto struct {
	Id        int64           `gorm:"column:id"`
	Balance   decimal.Decimal `gorm:"column:balance"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
}

func NewAccountDto(accEntity *entities.Account) *AccountDto {
	return &AccountDto{
		Id:      accEntity.Id,
		Balance: accEntity.Balance,
	}
}

func (dto *AccountDto) ToEntity() *entities.Account {
	return &entities.Account{
		Id:      dto.Id,
		Balance: dto.Balance,
	}
}

func (AccountDto) TableName() string {
	return "accounts"
}
