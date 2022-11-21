package helper

import (
	"go-rest/entity/domain"
	"go-rest/entity/web"
)

func ToHewanResponse(hewan domain.Hewan) web.HewanResponse {
	return web.HewanResponse{
		Id:   hewan.Id,
		Name: hewan.Name,
	}
}

func ToHewanResponses(semua_hewan []domain.Hewan) []web.HewanResponse {
	var hewanRes []web.HewanResponse
	for _, hewan := range semua_hewan {
		hewanRes = append(hewanRes, ToHewanResponse(hewan))
	}
	return hewanRes
}
