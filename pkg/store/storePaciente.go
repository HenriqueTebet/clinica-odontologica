package store

import (
	"clinica-odontologica/internal/domain"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	LoadEnvVariables()
}

type StorePaciente interface {
	Post(p domain.Paciente) (domain.Paciente, error)
	GetAll() []domain.Paciente
	GetById(id int) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

func NewStorePaciente() StorePaciente {
	dbPaciente, err := ConnectDb()

	if err != nil {
		panic(err)
	}

	return &storePaciente{
		db: dbPaciente,
	}
}

type storePaciente struct {
	db *sql.DB
}

func (sp *storePaciente) Post(p domain.Paciente) (domain.Paciente, error) {
	var paciente domain.Paciente

	result, err := sp.db.Exec("INSERT INTO pacientes (sobrenome, nome, registro_geral, data_cadastro) VALUES (?,?,?,?)", p.Sobrenome, p.Nome, p.RegistroGeral, p.DataCadastro)

	if err != nil {
		return paciente, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return paciente, err
	}

	paciente.Id = int(id)

	return sp.GetById(paciente.Id)
}

func (sp *storePaciente) GetAll() []domain.Paciente {
	var pacientes []domain.Paciente
	var p domain.Paciente

	result, err := sp.db.Query("SELECT * FROM pacientes")

	defer result.Close()

	if err != nil {
		return pacientes
	}

	for result.Next() {
		if err := result.Scan(
			&p.Id,
			&p.Sobrenome,
			&p.Nome,
			&p.RegistroGeral,
			&p.DataCadastro); err != nil {
			return pacientes
		}
		pacientes = append(pacientes, p)
	}

	return pacientes
}

func (sp *storePaciente) GetById(id int) (domain.Paciente, error) {
	var p domain.Paciente

	result, err := sp.db.Query("SELECT * FROM pacientes WHERE id = ?", id)

	defer result.Close()

	if err != nil {
		return domain.Paciente{}, err
	}

	defer result.Close()

	for result.Next() {
		if err = result.Scan(
			&p.Id,
			&p.Sobrenome,
			&p.Nome,
			&p.RegistroGeral,
			&p.DataCadastro); err != nil {
			return domain.Paciente{}, err
		}
		return p, nil
	}
	if result.Next() {
		return p, nil
	}
	return domain.Paciente{}, errors.New("Registro não encontrado")
}

func (sp *storePaciente) Update(id int, p domain.Paciente) (domain.Paciente, error) {

	_, err := sp.db.Exec("UPDATE pacientes SET sobrenome=?, nome=?, registro_geral=?, data_cadastro=? WHERE id=?", p.Sobrenome, p.Nome, p.RegistroGeral, p.DataCadastro, id)

	if err != nil {
		return p, err
	}

	return sp.GetById(id)
}

func (sp *storePaciente) Delete(id int) error {

	result, err := sp.db.Exec("DELETE FROM pacientes WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Paciente não encontrado")
	}

	if count != 0 {
		return nil
	}

	return err
}
