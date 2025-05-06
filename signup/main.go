package main

import (
	"autu-tutorial/controller"
	"autu-tutorial/internal/logger"
	"autu-tutorial/repository"
	"autu-tutorial/usecase"
	"log"
	"net/http"
)

var signUpController *controller.SignUpController

func init() {
	zerologger := logger.NewZeroLogger()
	establishedUserRepository := repository.NewEstablishedUserRepository(zerologger)
	signUpUsecase := usecase.NewSignUp(zerologger, establishedUserRepository)
	signUpController = controller.NewSignUpController(zerologger, signUpUsecase)
}

func main() {
	http.HandleFunc("/signup", signUpController.SignUpExec)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
