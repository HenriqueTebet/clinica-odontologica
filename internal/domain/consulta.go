package domain

type Consulta struct {
	Id          int     `json:"id"`
	Descricao        string  `json:"descricao" binding:"required"`
	DataHora  string  `json:"dataHora" binding:"required"`
	MatriculaDentista       string `json:"matriculaDentista" binding:"required"`
	RegistroPaciente	string `json:"registroPaciente" binding:"required"`
}