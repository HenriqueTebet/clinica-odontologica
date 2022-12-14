package dentista

import (
	"clinica-odontologica/internal/domain"
	"clinica-odontologica/pkg/store"
)

type Repository interface {
	Post(d domain.Dentista) (domain.Dentista, error)
	GetAll() []domain.Dentista
	GetById(id int) (domain.Dentista, error)
	Update(id int, d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	store store.StoreDentista
}

func NewRepository(store store.StoreDentista) Repository {
	return &repository{store}
}

func (r *repository) Post(d domain.Dentista) (domain.Dentista, error) {
	return r.store.Post(d)
}

func (r *repository) GetAll() []domain.Dentista {
	return r.store.GetAll()
}

func (r *repository) GetById(id int) (domain.Dentista, error) {
	return r.store.GetById(id)
}

func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	return r.store.Update(id, d)
}

func (r *repository) Delete(id int) error {
	return r.store.Delete(id)
}