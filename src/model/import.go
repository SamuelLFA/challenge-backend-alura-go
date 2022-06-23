package model

import (
	"time"
)

type ImportModel struct {
	Id                 int
	TimeOfImportation  time.Time
	TimeOfTransactions time.Time
}
