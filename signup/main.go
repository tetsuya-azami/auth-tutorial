package main

import (
	"autu-tutorial/controller"
	"autu-tutorial/internal/logger"
	"autu-tutorial/usecase"
	"log"
	"net/http"
)

var signUpController *controller.SignUpController

func init() {
	zerologger := logger.NewZeroLogger()
	signUpUsecase := usecase.NewSignUp(zerologger)
	signUpController = controller.NewSignUpController(zerologger, signUpUsecase)
}

func main() {
	http.HandleFunc("/signup", signUpController.SignUpExec)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
