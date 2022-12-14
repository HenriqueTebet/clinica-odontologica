package domain

type ConsultaDto struct {
	Consulta
	Dentista Dentista	`json:"dentist" binding:"required"`
	Paciente Paciente	`json:"patient" binding:"required"`
}