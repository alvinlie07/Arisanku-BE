package profile

import "gorm.io/gorm"

type Repository interface {
	BankList() (Bank, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) BankList() (Bank, error) {
	var bank Bank

	err := r.db.Debug().First(&bank).Error

	return bank, err
}
