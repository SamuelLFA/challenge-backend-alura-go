package model

import (
	"time"
)

const layoutISO = "2006-01-02"

type Import struct {
	TimeOfImportation  string
	TimeOfTransactions string
}

type ImportModel struct {
	Id                 int
	TimeOfImportation  time.Time
	TimeOfTransactions time.Time
}

func NewImport(importModel ImportModel) Import {
	return Import{
		TimeOfImportation:  importModel.TimeOfImportation.Format(layoutISO),
		TimeOfTransactions: importModel.TimeOfTransactions.Format(layoutISO),
	}
}
