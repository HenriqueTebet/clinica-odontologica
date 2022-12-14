package domain

type Dentista struct {
	Id          int     `json:"id"`
	Sobrenome        string  `json:"sobrenome" binding:"required"`
	Nome  string  `json:"nome" binding:"required"`
	Matricula       string `json:"matricula" binding:"required"`
}