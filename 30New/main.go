package main

import (
	"30New/controller"
	"30New/repo"
	"30New/usecase"
)

func main() {
	r := repo.New()
	useCase := usecase.New(r)
	serv := controller.New(useCase)
	controller.Server(serv)
}
