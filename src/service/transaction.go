package service

import (
	"challenge/src/model"
	"time"

	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func TransactionServiceFactory(db *gorm.DB) *TransactionService {
	return &TransactionService{db}
}

func (t *TransactionService) Save(transactions []model.Transaction) {
	var firstTimeTransaction time.Time
	firstTimeSetted := false
	for _, transaction := range transactions {
		model := transaction.ToModel()
		if !firstTimeSetted {
			firstTimeTransaction = model.Time
			firstTimeSetted = true
		}
		if model.Time.Equal(firstTimeTransaction) {
			t.db.Save(model)
		}
	}
}
