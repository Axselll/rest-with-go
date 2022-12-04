package controller

import (
	"go-rest/entity/web"
	"go-rest/helper"
	"go-rest/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type HewanControllerImpl struct {
	HewanService service.HewanService
}

func NewHewanController(hewanService service.HewanService) *HewanControllerImpl {
	return &HewanControllerImpl{
		HewanService: hewanService,
	}
}

func (controller *HewanControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	hewanCreateReq := web.HewanCreateReq{}
	helper.ReadFromReqBody(r, &hewanCreateReq)

	res := controller.HewanService.Create(r.Context(), hewanCreateReq)
	webRes := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *HewanControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	hewanUpdateReq := web.HewanUpdateReq{}
	helper.ReadFromReqBody(r, &hewanUpdateReq)

	hewan_id := params.ByName("id")
	id, err := strconv.Atoi(hewan_id)
	helper.CheckError(err)
	hewanUpdateReq.Id = id

	res := controller.HewanService.Update(r.Context(), hewanUpdateReq)
	webRes := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *HewanControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	hewan_id := params.ByName("id")
	id, err := strconv.Atoi(hewan_id)
	helper.CheckError(err)

	controller.HewanService.Delete(r.Context(), id)
	webRes := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *HewanControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	hewan_id := params.ByName("id")
	id, err := strconv.Atoi(hewan_id)
	helper.CheckError(err)

	res := controller.HewanService.FindById(r.Context(), id)
	webRes := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResBody(w, webRes)
}

func (controller *HewanControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := controller.HewanService.FindAll(r.Context())
	webRes := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	}

	helper.WriteToResBody(w, webRes)
}
