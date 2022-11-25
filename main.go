package main

import (
	"go-rest/app"
	"go-rest/controller"
	"go-rest/helper"
	"go-rest/middleware"
	"go-rest/repository"
	"go-rest/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load("local.env")
	helper.CheckError(errEnv)

	db := app.NewDB()
	validate := validator.New()
	hewanRepository := repository.NewHewanRepository()
	hewanService := service.NewHewanService(hewanRepository, db, validate)
	hewanController := controller.NewHewanController(hewanService)

	router := app.NewRouter(hewanController)

	server := http.Server{
		Addr:    ":6969",
		Handler: middleware.NewAuthMiddleware(router),
	}

	errServer := server.ListenAndServe()
	helper.CheckError(errServer)
}
