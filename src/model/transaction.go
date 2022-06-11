package model

import "time"

type transaction struct {
	originBank        string
	originAgency      string
	originAccount     string
	destinationBank   string
	destinatioAgency  string
	destinatioAccount string
	value             float64
	time              time.Time
}
