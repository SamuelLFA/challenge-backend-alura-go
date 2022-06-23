package service

import (
	m "challenge/src/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

func TransactionServiceFactory(db *gorm.DB) *TransactionService {
	return &TransactionService{db}
}

func (t *TransactionService) Save(transactions []m.Transaction) error {
	var firstTimeTransaction time.Time
	firstTimeSetted := false
	for _, transaction := range transactions {
		model := transaction.ToModel()
		if !firstTimeSetted {
			firstTimeTransaction = model.Time
			registerInTheSameDate := m.TransactionModel{}
			t.db.First(&registerInTheSameDate, m.TransactionModel{Time: firstTimeTransaction})
			if registerInTheSameDate.Id != "" {
				return errors.New(fmt.Sprintf("The transaction of %s was already uploaded", firstTimeTransaction))
			}
			firstTimeSetted = true
		}
		if model.Time.Equal(firstTimeTransaction) {
			t.db.Save(model)
		}
	}
	return nil
}
