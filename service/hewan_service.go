package service

import (
	"context"
	"go-rest/entity/web"
)

type HewanService interface {
	Create(ctx context.Context, req web.HewanCreateReq) web.HewanResponse
	Update(ctx context.Context, req web.HewanUpdateReq) web.HewanResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.HewanResponse
	FindAll(ctx context.Context) []web.HewanResponse
}
