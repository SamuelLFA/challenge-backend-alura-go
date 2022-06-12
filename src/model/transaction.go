package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	originBank        string
	originAgency      string
	originAccount     string
	destinationBank   string
	destinatioAgency  string
	destinatioAccount string
	value             float64
	time              time.Time
}

type TransactionModel struct {
	Id                string    `gorm:"primaryKey"`
	OriginBank        string    `gorm:"notNull"`
	OriginAgency      string    `gorm:"notNull"`
	OriginAccount     string    `gorm:"notNull"`
	DestinationBank   string    `gorm:"notNull"`
	DestinatioAgency  string    `gorm:"notNull"`
	DestinatioAccount string    `gorm:"notNull"`
	Value             float64   `gorm:"notNull"`
	Time              time.Time `gorm:"notNull"`
}

func (t *TransactionModel) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.NewString()
	return
}

func NewTransaction(
	originBank string,
	originAgency string,
	originAccount string,
	destinationBank string,
	destinatioAgency string,
	destinatioAccount string,
	value float64,
	time time.Time) *Transaction {
	return &Transaction{
		originBank,
		originAgency,
		originAccount,
		destinationBank,
		destinatioAgency,
		destinatioAccount,
		value,
		time,
	}
}

func (t *Transaction) ToModel() *TransactionModel {
	return &TransactionModel{
		OriginBank:        t.originBank,
		OriginAgency:      t.originAgency,
		OriginAccount:     t.originAccount,
		DestinationBank:   t.destinationBank,
		DestinatioAgency:  t.destinatioAgency,
		DestinatioAccount: t.destinatioAccount,
		Value:             t.value,
		Time:              t.time,
	}
}
