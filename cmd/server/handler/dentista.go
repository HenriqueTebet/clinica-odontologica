package handler

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/internal/dentista"
	"clinica-odontologica/pkg/web"
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{s:s}
}

func (ph *dentistaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentista domain.Dentista

		err := ctx.ShouldBindJSON(&dentista)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmptyDentista(&dentista)

		if isValid {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Post(dentista)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (ph *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := ph.s.GetAll()

		if response == nil {
			web.Failure(ctx, 404, "Not Found", "Não foi encontrado nenhum registro")
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *dentistaHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		response, err := ph.s.GetById(id)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", "Não foi encontrado nenhum registro")
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *dentistaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		var dentista domain.Dentista

		err = ctx.ShouldBindJSON(&dentista)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmptyDentista(&dentista)

		if isValid {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Update(id, dentista)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *dentistaHandler) Patch() gin.HandlerFunc {
	type request struct {
		Sobrenome        string  `json:"sobrenome,omitempty"`
		Nome  string  `json:"nome,omitempty"`
		Matricula       string `json:"matricula,omitempty"`
	}

	return func(ctx *gin.Context) {

		var request request

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		err = ctx.ShouldBindJSON(&request)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		updateDentista := domain.Dentista{
			Sobrenome: request.Sobrenome,
			Nome: request.Nome,
			Matricula: request.Matricula,
		}

		response, err := ph.s.Update(id, updateDentista)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *dentistaHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		err = ph.s.Delete(id)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		ctx.JSON(204, "")
	}
}

func isEmptyDentista(d *domain.Dentista) (bool, error) {

	if d.Sobrenome != ""|| d.Nome != ""|| d.Matricula != "" {
		return false, errors.New("Todos os campos devem ser preenchidos")
	}

	return true, nil
}