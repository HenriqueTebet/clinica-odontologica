package handler

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/internal/consulta"
	"clinica-odontologica/pkg/web"
	"errors"
	"strconv"
	"github.com/gin-gonic/gin"
)

type consultaHandler struct {
	s consulta.Service
}

func NewConsultaHandler(s consulta.Service) *consultaHandler {
	return &consultaHandler{s:s}
}

func (ph *consultaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var consulta domain.Consulta

		err := ctx.ShouldBindJSON(&consulta)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmptyConsulta(&consulta)

		if isValid {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Post(consulta)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (ph *consultaHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := ph.s.GetAll()

		if response == nil {
			web.Failure(ctx, 404, "Not Found", "Não foi encontrado nenhum registro")
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *consultaHandler) GetById() gin.HandlerFunc {
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

func (ph *consultaHandler) GetByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rgParam := ctx.Param("rg")

		_, err := strconv.Atoi(rgParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "RG Inválido")
			return
		}

		response, err := ph.s.GetByRg(rgParam)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", "Não foi encontrado nenhum registro")
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *consultaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "ID Inválido")
			return
		}

		var consulta domain.Consulta

		err = ctx.ShouldBindJSON(&consulta)

		if err != nil {
			web.Failure(ctx, 400, "Bad Request", "Parâmetros inválidos")
			return
		}

		isValid, err := isEmptyConsulta(&consulta)

		if isValid {
			web.Failure(ctx, 400, "Bad Request", err.Error())
			return
		}

		response, err := ph.s.Update(id, consulta)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *consultaHandler) Patch() gin.HandlerFunc {
	type request struct {
		Descricao        string  `json:"descricao,omitempty"`
		DataHora  string  `json:"dataHora,omitempty"`
		MatriculaDentista       string `json:"matriculaDentista,omitempty"`
		RegistroPaciente	string `json:"registroPaciente,omitempty"`
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

		updateConsulta := domain.Consulta{
			Descricao: request.Descricao,
			DataHora: request.DataHora,
			MatriculaDentista: request.MatriculaDentista,
			RegistroPaciente: request.RegistroPaciente,
		}

		response, err := ph.s.Update(id, updateConsulta)

		if err != nil {
			web.Failure(ctx, 404, "Not Found", err.Error())
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *consultaHandler) Delete() gin.HandlerFunc {
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

func isEmptyConsulta(c *domain.Consulta) (bool, error) {

	if c.Descricao != ""|| c.DataHora != ""|| c.MatriculaDentista != "" || c.RegistroPaciente != "" {
		return false, errors.New("Todos os campos devem ser preenchidos")
	}

	return true, nil
}