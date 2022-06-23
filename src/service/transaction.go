package service

import (
	m "challenge/src/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const layoutISO = "2006-01-02"

type TransactionService struct {
	db            *gorm.DB
	importService *ImportService
}

func TransactionServiceFactory(db *gorm.DB, importService *ImportService) *TransactionService {
	return &TransactionService{db, importService}
}

func (t *TransactionService) Save(transactions []m.Transaction) error {
	var firstTimeTransaction time.Time
	firstTimeSetted := false
	for _, transaction := range transactions {
		model := transaction.ToModel()
		if !firstTimeSetted {
			firstTimeTransaction = model.Time
			alreadyImported := t.importService.AnyImportAt(firstTimeTransaction)
			if alreadyImported {
				return errors.New(fmt.Sprintf("The transaction of %s was already uploaded", firstTimeTransaction))
			}
			importModel := m.ImportModel{TimeOfImportation: time.Now(), TimeOfTransactions: firstTimeTransaction}
			t.importService.Save(&importModel)
			firstTimeSetted = true
		}
		if model.Time.Format(layoutISO) == firstTimeTransaction.Format(layoutISO) {
			t.db.Save(model)
		}
	}
	return nil
}
