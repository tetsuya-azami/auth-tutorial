package controller

import (
	"autu-tutorial/domain"
	"autu-tutorial/internal/logger"
	"autu-tutorial/usecase"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SignUp interface {
	Execute(email domain.Email, password domain.Password)
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpController struct {
	signUp SignUp
	logger logger.Logger
}

func NewSignUpController(l logger.Logger, su *usecase.SignUp) *SignUpController {
	return &SignUpController{
		logger: l,
		signUp: su,
	}
}

func (suc *SignUpController) SignUpExec(w http.ResponseWriter, r *http.Request) {
	suc.logger.LogConnectionInfo(r)

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed")
		return
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BadRequest, ", err)
		return
	}

	var request SignUpRequest
	json.Unmarshal(body, &request)

	email, err := domain.NewEmail(request.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	password, err := domain.NewPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	suc.signUp.Execute(email, password)

	fmt.Fprint(w, "received, ")
}
