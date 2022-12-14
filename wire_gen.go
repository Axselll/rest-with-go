// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go-rest/app"
	"go-rest/controller"
	"go-rest/middleware"
	"go-rest/repository"
	"go-rest/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitServer() *http.Server {
	hewanRepositoryImpl := repository.NewHewanRepository()
	db := app.NewDB()
	validate := validator.New()
	hewanServiceImpl := service.NewHewanService(hewanRepositoryImpl, db, validate)
	hewanControllerImpl := controller.NewHewanController(hewanServiceImpl)
	router := app.NewRouter(hewanControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var hewanSet = wire.NewSet(repository.NewHewanRepository, wire.Bind(new(repository.HewanRepository), new(*repository.HewanRepositoryImpl)), service.NewHewanService, wire.Bind(new(service.HewanService), new(*service.HewanServiceImpl)), controller.NewHewanController, wire.Bind(new(controller.HewanController), new(*controller.HewanControllerImpl)))
