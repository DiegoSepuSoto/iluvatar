package requests

type HookRequest struct {
	Entry interface{} `json:"entry"`
	Event string `json:"event"`
	Model string `json:"model"`
}

type HookPostModelRequest struct {
	ID string `json:"id"`
	Title string `json:"Titulo"`
	Overview string `json:"Resumen"`
	Service *Service `json:"servicio"`
}

type Service struct {
	Name string `json:"Nombre"`
	ShortName string `json:"Sigla"`
}