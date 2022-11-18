package repository

import (
	"context"
	"database/sql"
	"go-rest/entity/domain"
)

type HewanRepository interface {
	Save(ctx context.Context, tx *sql.Tx, hewan domain.Hewan) domain.Hewan
	Update(ctx context.Context, tx *sql.Tx, hewan domain.Hewan) domain.Hewan
	Delete(ctx context.Context, tx *sql.Tx, hewan domain.Hewan)
	FindByID(ctx context.Context, tx *sql.Tx, id int) (domain.Hewan, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Hewan
}
