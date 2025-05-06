package domain

import (
	"errors"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	// 本来はRFCに準拠すべきだが簡易的に作成
	if len(value) == 0 {
		return Email{}, errors.New("email must not be empty")
	}
	if !strings.Contains(value, "@") {
		return Email{}, errors.New("email must have @")
	}

	return Email{value: value}, nil
}

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if len(value) < 8 {
		return Password{}, errors.New("password must consist of at least 8 characters")
	}
	return Password{value: value}, nil
}
