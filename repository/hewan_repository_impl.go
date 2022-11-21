package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest/entity/domain"
	"go-rest/helper"
)

type HewanRepositoryImpl struct {
}

func NewHewanRepository() HewanRepository {
	return &HewanRepositoryImpl{}
}

func (repository *HewanRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, hewan domain.Hewan) domain.Hewan {
	query := "insert into hewan(name) values (?)"

	res, err := tx.ExecContext(ctx, query, hewan.Name)
	helper.CheckError(err)

	id, err := res.LastInsertId()
	helper.CheckError(err)

	hewan.Id = int(id)
	return hewan
}

func (repository *HewanRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, hewan domain.Hewan) domain.Hewan {
	query := "update hewan set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, hewan.Name, hewan.Id)
	helper.CheckError(err)

	return hewan
}

func (repository *HewanRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, hewan domain.Hewan) {
	query := "delete from hewan where id = ?"
	_, err := tx.ExecContext(ctx, query, hewan.Id)
	helper.CheckError(err)
}

func (repository *HewanRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, id int) (domain.Hewan, error) {
	query := "select id, name from hewan where id = ?"
	res, err := tx.QueryContext(ctx, query, id)
	helper.CheckError(err)
	defer res.Close()

	hewan := domain.Hewan{}
	if res.Next() {
		err := res.Scan(&hewan.Id, &hewan.Name)
		helper.CheckError(err)
		return hewan, nil
	} else {
		return hewan, errors.New("hewan not found")
	}
}

func (repository *HewanRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Hewan {
	query := "select id, name from hewan"
	res, err := tx.QueryContext(ctx, query)
	helper.CheckError(err)
	defer res.Close()

	var semua_hewan []domain.Hewan
	for res.Next() {
		hewan := domain.Hewan{}
		err := res.Scan(&hewan.Id, &hewan.Name)
		helper.CheckError(err)
		semua_hewan = append(semua_hewan, hewan)
	}
	return semua_hewan
}
