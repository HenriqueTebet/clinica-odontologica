package store

import (
	"clinica-odontologica/internal/domain"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type StoreDentista interface {
	Post(p domain.Dentista) (domain.Dentista, error)
	GetAll() []domain.Dentista
	GetById(id int) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

func NewStoreDentista() StoreDentista {
	dbDentista, err := ConnectDb()

	if err != nil {
		panic(err)
	}

	return &storeDentista{
		db: dbDentista,
	}
}

type storeDentista struct {
	db *sql.DB
}

func (sd *storeDentista) Post(d domain.Dentista) (domain.Dentista, error) {
	var dentista domain.Dentista

	result, err := sd.db.Exec("INSERT INTO dentistas (sobrenome, nome, matricula) VALUES (?,?,?)", d.Sobrenome, d.Nome, d.Matricula)

	if err != nil {
		return dentista, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return dentista, err
	}

	dentista.Id = int(id)

	return sd.GetById(dentista.Id)
}

func (sd *storeDentista) GetAll() []domain.Dentista {
	var dentistas []domain.Dentista
	var d domain.Dentista

	result, err := sd.db.Query("SELECT * FROM dentistas")

	defer result.Close()

	if err != nil {
		return dentistas
	}

	for result.Next() {
		if err := result.Scan(
			&d.Id,
			&d.Sobrenome,
			&d.Nome,
			&d.Matricula); err != nil {
				return dentistas
			}
		dentistas = append(dentistas, d)
	}

	return dentistas
}

func (sd *storeDentista) GetById(id int) (domain.Dentista, error) {
	var d domain.Dentista

	result, err := sd.db.Query("SELECT * FROM dentistas WHERE id = ?", id)

	defer result.Close()

	if err != nil {
		return d, err
	}

	for result.Next() {
		if err := result.Scan(
			&d.Id,
			&d.Sobrenome,
			&d.Nome,
			&d.Matricula); err != nil {
				return d, err
			}

			return d, nil
	}

	if result.Next() {
		return d, nil
	}

	return domain.Dentista{}, errors.New("Registro não encontrado")
}

func (sd *storeDentista) Update(id int, d domain.Dentista) (domain.Dentista, error) {

	_, err := sd.db.Exec("UPDATE dentistas SET sobrenome=?, nome=?, matricula=? WHERE id=?", d.Sobrenome, d.Nome, d.Matricula, id)

	if err != nil {
		return d, err
	}

	return sd.GetById(id)
}

func (sd *storeDentista) Delete(id int) error {

	result, err := sd.db.Exec("DELETE FROM dentistas WHERE id = ?", id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return errors.New("Dentista não encontrado")
	}

	if count != 0 {
		return nil
	}

	return err
}