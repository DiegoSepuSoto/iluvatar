package entity

type LoginEntity struct {
	Sesion         string    `json:"sesion"`
	TiposUsuario   []string `json:"tiposUsuario"`
	Rut            int       `json:"rut"`
	NombreCompleto string    `json:"nombreCompleto"`
	FotoURL        string    `json:"fotoUrl"`
	Correo         string    `json:"correo"`
}
