package service

import (
	"context"
	"database/sql"
	"go-rest/entity/domain"
	"go-rest/entity/web"
	"go-rest/exception"
	"go-rest/helper"
	"go-rest/repository"

	"github.com/go-playground/validator/v10"
)

type HewanServiceImpl struct {
	HewanRepository repository.HewanRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewHewanService(hewanRepository repository.HewanRepository, DB *sql.DB, Validate *validator.Validate) HewanService {
	return &HewanServiceImpl{
		HewanRepository: hewanRepository,
		DB:              DB,
		Validate:        Validate}
}

func (service *HewanServiceImpl) Create(ctx context.Context, req web.HewanCreateReq) web.HewanResponse {
	err := service.Validate.Struct(req)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	hewan := domain.Hewan{
		Name: req.Name,
	}

	hewan = service.HewanRepository.Save(ctx, tx, hewan)

	return helper.ToHewanResponse(hewan)
}

func (service *HewanServiceImpl) Update(ctx context.Context, req web.HewanUpdateReq) web.HewanResponse {
	err := service.Validate.Struct(req)
	helper.CheckError(err)

	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	hewan, err := service.HewanRepository.FindByID(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	hewan.Name = req.Name

	hewan = service.HewanRepository.Update(ctx, tx, hewan)

	return helper.ToHewanResponse(hewan)
}

func (service *HewanServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	hewan, err := service.HewanRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	service.HewanRepository.Delete(ctx, tx, hewan)
}

func (service *HewanServiceImpl) FindById(ctx context.Context, id int) web.HewanResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	hewan, err := service.HewanRepository.FindByID(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundErr(err.Error()))
	}

	return helper.ToHewanResponse(hewan)
}

func (service *HewanServiceImpl) FindAll(ctx context.Context) []web.HewanResponse {
	tx, err := service.DB.Begin()
	helper.CheckError(err)
	defer helper.CommitOrRollback(tx)

	semua_hewan := service.HewanRepository.FindAll(ctx, tx)

	return helper.ToHewanResponses(semua_hewan)
}
