package app

import (
	"go-rest/controller"
	"go-rest/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(hewanController controller.HewanController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/hewan", hewanController.FindAll)
	router.GET("/api/hewan/:id", hewanController.FindById)
	router.POST("/api/hewan", hewanController.Create)
	router.PUT("/api/hewan/:id", hewanController.Update)
	router.DELETE("/api/hewan/:id", hewanController.Delete)

	router.PanicHandler = exception.ErrHandler

	return router
}
