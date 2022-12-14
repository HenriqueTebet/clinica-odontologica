package domain

type Paciente struct {
	Id          int     `json:"id"`
	Sobrenome        string  `json:"sobrenome" binding:"required"`
	Nome  string  `json:"nome" binding:"required"`
	RegistroGeral       string `json:"registroGeral" binding:"required"`
	DataCadastro	string `json:"dataCadastro" binding:"required"`
}