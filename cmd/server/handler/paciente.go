package handler

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/internal/paciente"
	"clinica-odontologica/pkg/web"
	"errors"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{s:s}
}

func (ph *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente

		err := ctx.ShouldBindJSON(&paciente)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmpty(&paciente)
		validateFormat, err := dateFormat(paciente.DataCadastro)

		if isValid || validateFormat {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Post(paciente)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (ph *pacienteHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := ph.s.GetAll()

		if response == nil {
			web.Failure(ctx, 404, "Not Found", "Não foi encontrado nenhum registro")
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) GetById() gin.HandlerFunc {
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

func (ph *pacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		var paciente domain.Paciente

		err = ctx.ShouldBindJSON(&paciente)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmpty(&paciente)
		validateFormat, err := dateFormat(paciente.DataCadastro)

		if isValid || validateFormat {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Update(id, paciente)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) Patch() gin.HandlerFunc {
	type request struct {
		Sobrenome        string  `json:"sobrenome,omitempty"`
		Nome  string  `json:"nome,omitempty"`
		RegistroGeral       string `json:"registroGeral,omitempty"`
		DataCadastro	string `json:"dataCadastro,omitempty"`
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

		updatePaciente := domain.Paciente{
			Sobrenome: request.Sobrenome,
			Nome: request.Nome,
			RegistroGeral: request.RegistroGeral,
			DataCadastro: request.DataCadastro,
		}

		validateFormat, err := dateFormat(request.DataCadastro)

		if validateFormat {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Update(id, updatePaciente)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) Delete() gin.HandlerFunc {
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

func isEmpty(p *domain.Paciente) (bool, error) {

	if p.Sobrenome != ""|| p.Nome != ""|| p.RegistroGeral != ""|| p.DataCadastro != "" {
		return false, errors.New("Todos os campos devem ser preenchidos")
	}

	return true, nil
}

func dateFormat(dh string) (bool, error) {
	_, err := time.Parse("dd/MM/yyyy hh:mm", dh)

	if err != nil {
		return false, errors.New("Data inválida")
	}

	return true, nil
}