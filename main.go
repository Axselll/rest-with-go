package main

import (
	"go-rest/app"
	"go-rest/controller"
	"go-rest/exception"
	"go-rest/helper"
	"go-rest/repository"
	"go-rest/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	hewanRepository := repository.NewHewanRepository()
	hewanService := service.NewHewanService(hewanRepository, db, validate)
	hewanContoller := controller.NewHewanController(hewanService)

	router := httprouter.New()

	router.GET("/api/hewan", hewanContoller.FindAll)
	router.GET("/api/hewan/:id", hewanContoller.FindById)
	router.POST("/api/hewan", hewanContoller.Create)
	router.PUT("/api/hewan/:id", hewanContoller.Update)
	router.DELETE("/api/hewan/:id", hewanContoller.Delete)

	router.PanicHandler = exception.ErrHandler

	server := http.Server{
		Addr:    "localhost:6969",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.CheckError(err)
}
