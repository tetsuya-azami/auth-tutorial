package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
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

func (e Email) Value() string {
	return e.value
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

func (p Password) Value() string {
	return p.value
}

type SignUpToken struct {
	value string
}

func NewSignUpToken() (SignUpToken, error) {
	value, err := uuid.NewRandom()
	if err != nil {
		return SignUpToken{}, err
	}

	return SignUpToken{value: value.String()}, nil
}

func (sut SignUpToken) Value() string {
	return sut.value
}
