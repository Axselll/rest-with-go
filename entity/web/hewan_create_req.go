package web

type HewanCreateReq struct {
	Name string `json:"name" validate:"required,max=50,min=1"`
}
