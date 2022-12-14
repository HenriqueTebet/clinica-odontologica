package main

import (
	"clinica-odontologica/internal/consulta"
	"clinica-odontologica/internal/dentista"
	"clinica-odontologica/internal/paciente"
	"clinica-odontologica/pkg/store"
	"clinica-odontologica/cmd/server/handler"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err.Error())
	}

	sqlConsulta := store.NewStoreConsulta()
	sqlDentista := store.NewStoreDentista()
	sqlPaciente := store.NewStorePaciente()

	repoConsulta := consulta.NewRepository(sqlConsulta)
	repoDentista := dentista.NewRepository(sqlDentista)
	repoPaciente := paciente.NewRepository(sqlPaciente)

	serviceConsulta := consulta.NewService(repoConsulta)
	serviceDentista := dentista.NewService(repoDentista)
	servicePaciente := paciente.NewService(repoPaciente)

	handlerConsulta := handler.NewConsultaHandler(serviceConsulta)
	handlerDentista := handler.NewDentistaHandler(serviceDentista)
	handlerPaciente := handler.NewPacienteHandler(servicePaciente)

	app := gin.New()

	app.Use(gin.Recovery(), gin.Logger())

	consultas := app.Group("/consultas")
	{
		consultas.POST("", handlerConsulta.Post())
		consultas.GET("", handlerConsulta.GetAll())
		consultas.GET(":id", handlerConsulta.GetById())
		consultas.GET("paciente/:rg", handlerConsulta.GetByRg())
		consultas.PUT(":id", handlerConsulta.Put())
		consultas.PATCH(":id", handlerConsulta.Patch())
		consultas.DELETE(":id", handlerConsulta.Delete())
	}

	dentistas := app.Group("/dentistas")
	{
		dentistas.POST("", handlerDentista.Post())
		dentistas.GET("", handlerDentista.GetAll())
		dentistas.GET(":id", handlerDentista.GetById())
		dentistas.PUT(":id", handlerDentista.Put())
		dentistas.PATCH(":id", handlerDentista.Patch())
		dentistas.DELETE(":id", handlerDentista.Delete())
	}

	pacientes := app.Group("/pacientes")
	{
		pacientes.POST("", handlerPaciente.Post())
		pacientes.GET("", handlerPaciente.GetAll())
		pacientes.GET(":id", handlerPaciente.GetById())
		pacientes.PUT(":id", handlerPaciente.Put())
		pacientes.PATCH(":id", handlerPaciente.Patch())
		pacientes.DELETE(":id", handlerPaciente.Delete())
	}

	app.Run(":8080")
}