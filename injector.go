//go:build wireinject
// +build wireinject

package main

import (
	"go-rest/app"
	"go-rest/controller"
	"go-rest/middleware"
	"go-rest/repository"
	"go-rest/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var hewanSet = wire.NewSet(
	repository.NewHewanRepository,
	wire.Bind(new(repository.HewanRepository), new(*repository.HewanRepositoryImpl)),
	service.NewHewanService,
	wire.Bind(new(service.HewanService), new(*service.HewanServiceImpl)),
	controller.NewHewanController,
	wire.Bind(new(controller.HewanController), new(*controller.HewanControllerImpl)),
)

func InitServer() *http.Server {
	wire.Build(app.NewDB, validator.New, hewanSet, app.NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)), middleware.NewAuthMiddleware, NewServer)
	return nil
}
