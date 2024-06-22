package repositories

import (
	"com/txfer/bounded_contexts/account/dtos"
	"com/txfer/bounded_contexts/account/entities"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type AccountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (repo *AccountRepo) CreateAccount(account *entities.Account) error {
	dto := dtos.NewAccountDto(account)
	if err := repo.db.Create(dto).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return errors.New("account id already exists")
		}
		return err
	}

	return nil
}

func (repo *AccountRepo) GetAccountById(id int64) (*entities.Account, error) {
	var dto dtos.AccountDto
	if err := repo.db.Where("id = ?", id).First(&dto).Error; err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil
}
