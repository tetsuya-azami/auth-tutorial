package repository

import (
	"autu-tutorial/domain"
	"autu-tutorial/internal/logger"
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type EstablishedUserRepository struct {
	logger logger.Logger
}

func NewEstablishedUserRepository(logger logger.Logger) *EstablishedUserRepository {
	return &EstablishedUserRepository{
		logger: logger,
	}
}

func (ur *EstablishedUserRepository) Save(email domain.Email, password domain.Password, token domain.SignUpToken) error {
	db, err := sql.Open("mysql", "root:password@/auth_tutorial")
	if err != nil {
		ur.logger.LogError(err)
		return err
	}

	_, err = db.ExecContext(
		context.Background(),
		"INSERT INTO establised_users (email, password, token) VALUES($1, $2, $3)",
		email.Value(),
		password.Value(),
		token.Value(),
	)

	return err
}
