package consulta

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/pkg/store"
)

type Repository interface {
	Post(c domain.Consulta) (domain.ConsultaDto, error)
	GetAll() []domain.ConsultaDto
	GetById(id int) (domain.ConsultaDto, error)
	GetByRg(rg string) ([]domain.ConsultaDto, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDto, error)
	Delete(id int) error
}

type repository struct {
	store store.StoreConsulta
}

func NewRepository(store store.StoreConsulta) Repository {
	return &repository{store}
}

func (r *repository) Post(c domain.Consulta) (domain.ConsultaDto, error) {
	return r.store.Post(c)
}

func (r *repository) GetAll() []domain.ConsultaDto {
	return r.store.GetAll()
}

func (r *repository) GetById(id int) (domain.ConsultaDto, error) {
	return r.store.GetById(id)
}

func (r *repository) GetByRg(rg string) ([]domain.ConsultaDto, error) {
	return r.store.GetByRg(rg)
}

func (r *repository) Update(id int, c domain.Consulta) (domain.ConsultaDto, error) {
	return r.store.Update(id, c)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}