package service

import (
	"challenge/src/model"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type FileService struct{}

func FileServiceFactory() *FileService {
	return &FileService{}
}

func (f *FileService) ParseLines(text string, transactions *[]model.Transaction) error {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		transaction, err := stringToTransaction(line)
		if err != nil {
			return err
		}
		*transactions = append(*transactions, *transaction)
	}
	return nil
}

func stringToTransaction(line string) (*model.Transaction, error) {
	fields := strings.Split(line, ",")
	if len(fields) < 8 {
		log.Printf("Line with only %d field\n", len(fields))
		return nil, errors.New("Failed to parse CSV")
	}

	date, err := time.Parse("2006-01-02T15:04:05", fields[7])

	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return nil, err
	}

	return model.NewTransaction(
		fields[0],
		fields[1],
		fields[2],
		fields[3],
		fields[4],
		fields[5],
		value,
		date,
	), nil
}
