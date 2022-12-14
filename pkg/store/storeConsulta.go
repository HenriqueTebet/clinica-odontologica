package store

import (
	"clinica-odontologica/internal/domain"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type StoreConsulta interface {
	Post(c domain.Consulta) (domain.ConsultaDto, error)
	GetAll() []domain.ConsultaDto
	GetById(id int) (domain.ConsultaDto, error)
	GetByRg(rg string) ([]domain.ConsultaDto, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDto, error)
	Delete(id int) error
}

func NewStoreConsulta() StoreConsulta {
	dbConsulta, err := ConnectDb()

	if err != nil {
		panic(err)
	}

	return &storeConsulta{
		db: dbConsulta,
	}
}

type storeConsulta struct {
	db *sql.DB
}

func (sc *storeConsulta) Post(c domain.Consulta) (domain.ConsultaDto, error) {
	result, err := sc.db.Exec("INSERT INTO consultas (descricao, data_hora, matricula_dentista, registro_paciente) VALUES (?,?,?,?)", c.Descricao, c.DataHora, c.MatriculaDentista, c.RegistroPaciente)

	if err != nil {
		return domain.ConsultaDto{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return domain.ConsultaDto{}, err
	}

	return sc.GetById(int(id))
}

func (sc *storeConsulta) GetAll() []domain.ConsultaDto {
	var consultas []domain.ConsultaDto
	var c domain.ConsultaDto

	result, err := sc.db.Query("SELECT * FROM consultas c inner JOIN dentistas d on c.matricula_dentista = d.matricula inner JOIN pacientes p on c.registro_paciente = p.registro_geral")

	defer result.Close()

	if err != nil {
		return consultas
	}

	for result.Next() {
		if err := result.Scan(
			&c.Id,
			&c.Descricao,
			&c.DataHora,
			&c.MatriculaDentista,
			&c.RegistroPaciente,
			&c.Dentista.Id,
			&c.Dentista.Sobrenome,
			&c.Dentista.Nome,
			&c.Dentista.Matricula,
			&c.Paciente.Id,
			&c.Paciente.Sobrenome,
			&c.Paciente.Nome,
			&c.Paciente.RegistroGeral,
			&c.Paciente.DataCadastro); err != nil {
			return consultas
		}
		consultas = append(consultas, c)
	}

	return consultas
}

func (sc *storeConsulta) GetById(id int) (domain.ConsultaDto, error) {
	var c domain.ConsultaDto

	result, err := sc.db.Query("SELECT * FROM consultas c inner JOIN dentistas d on c.matricula_dentista = d.matricula inner JOIN pacientes p on c.registro_paciente = p.registro_geral WHERE c.id=?", id)

	defer result.Close()

	if err != nil {
		return c, err
	}

	for result.Next() {
		if err := result.Scan(
			&c.Id,
			&c.Descricao,
			&c.DataHora,
			&c.MatriculaDentista,
			&c.RegistroPaciente,
			&c.Dentista.Id,
			&c.Dentista.Sobrenome,
			&c.Dentista.Nome,
			&c.Dentista.Matricula,
			&c.Paciente.Id,
			&c.Paciente.Sobrenome,
			&c.Paciente.Nome,
			&c.Paciente.RegistroGeral,
			&c.Paciente.DataCadastro); err != nil {
			return c, err
		}

		return c, nil
	}

	if result.Next() {
		return c, nil
	}

	return domain.ConsultaDto{}, errors.New("Registro não encontrado")
}

func (sc *storeConsulta) GetByRg(rg string) ([]domain.ConsultaDto, error) {
	var c domain.ConsultaDto
	var consultas []domain.ConsultaDto

	result, err := sc.db.Query("SELECT * FROM consultas c inner JOIN dentistas d on c.matricula_dentista = d.matricula inner JOIN pacientes p on c.registro_paciente = p.registro_geral WHERE c.registro_paciente=?", rg)

	defer result.Close()

	if err != nil {
		return consultas, err
	}

	for result.Next() {
		if err := result.Scan(
			&c.Id,
			&c.Descricao,
			&c.DataHora,
			&c.MatriculaDentista,
			&c.RegistroPaciente,
			&c.Dentista.Id,
			&c.Dentista.Sobrenome,
			&c.Dentista.Nome,
			&c.Dentista.Matricula,
			&c.Paciente.Id,
			&c.Paciente.Sobrenome,
			&c.Paciente.Nome,
			&c.Paciente.RegistroGeral,
			&c.Paciente.DataCadastro); err != nil {
			return consultas, err
		}
		consultas = append(consultas, c)
	}

	return consultas, nil
}

func (sc *storeConsulta) Update(id int, c domain.Consulta) (domain.ConsultaDto, error) {

	_, err := sc.db.Exec("UPDATE consultas SET descricao=?, data_hora=?, matricula_dentista=?, registro_paciente=? WHERE id=?", c.Descricao, c.DataHora, c.MatriculaDentista, c.RegistroPaciente, id)

	if err != nil {
		return domain.ConsultaDto{}, err
	}

	return sc.GetById(id)
}

func (sc *storeConsulta) Delete(id int) error {

	result, err := sc.db.Exec("DELETE FROM consultas WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Consulta não encontrado")
	}

	if count != 0 {
		return nil
	}

	return err
}
