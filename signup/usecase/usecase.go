package usecase

import (
	"autu-tutorial/domain"
	"autu-tutorial/internal/logger"
)

type SignUp struct {
	logger                    logger.Logger
	establishedUserRepository EstablishedUserRepository
}

type EstablishedUserRepository interface {
	Save(email domain.Email, password domain.Password, token domain.SignUpToken) error
}

func NewSignUp(l logger.Logger, eur EstablishedUserRepository) *SignUp {
	return &SignUp{
		logger:                    l,
		establishedUserRepository: eur,
	}
}

func (su *SignUp) Execute(email domain.Email, password domain.Password) {
	token, err := domain.NewSignUpToken()
	if err != nil {

	}
	// DBへの登録
	su.establishedUserRepository.Save(email, password, token)

	// メール送信
}
