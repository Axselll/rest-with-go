package web

type HewanUpdateReq struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=30,min=1"`
}
