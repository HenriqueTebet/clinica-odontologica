package paciente

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/pkg/store"
)

type Repository interface {
	Post(p domain.Paciente) (domain.Paciente, error)
	GetAll() []domain.Paciente
	GetById(id int) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	store store.StorePaciente
}

func NewRepository(store store.StorePaciente) Repository {
	return &repository{store}
}

func (r *repository) Post(p domain.Paciente) (domain.Paciente, error) {
	return r.store.Post(p)
}

func (r *repository) GetAll() []domain.Paciente {
	return r.store.GetAll()
}

func (r *repository) GetById(id int) (domain.Paciente, error) {
	return r.store.GetById(id)
}

func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	return r.store.Update(id, p)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}