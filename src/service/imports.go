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

func (s *ImportService) FindAll() *[]m.Import {
	models := []m.ImportModel{}
	s.db.Find(&models)
	imports := []m.Import{}

	for _, model := range models {
		imports = append(imports, m.NewImport(model))
	}
	return &imports
}

func (s *ImportService) AnyImportAt(time time.Time) bool {
	model := model.ImportModel{}
	s.db.First(&model, m.ImportModel{TimeOfTransactions: time})
	return model.Id != 0
}
