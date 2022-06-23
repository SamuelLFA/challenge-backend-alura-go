package service

import (
	"challenge/src/model"
	m "challenge/src/model"
	"time"

	"gorm.io/gorm"
)

type ImportService struct {
	db *gorm.DB
}

func ImportServiceFactory(db *gorm.DB) *ImportService {
	return &ImportService{
		db: db,
	}
}

func (s *ImportService) Save(importModel *m.ImportModel) {
	s.db.Save(importModel)
}

func (s *ImportService) AnyImportAt(time time.Time) bool {
	model := model.ImportModel{}
	s.db.First(&model, m.ImportModel{TimeOfTransactions: time})
	return model.Id != 0
}
