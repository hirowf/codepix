package model

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/martian/log"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"-"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		log.Errorf("error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	err = t.isValid()
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransction() *Transaction {
	return &Transaction{}
}
