package consulta

import (
	"clinica-odontologica/internal/domain"
)

type Service interface {
	Post(c domain.Consulta) (domain.ConsultaDto, error)
	GetAll() []domain.ConsultaDto
	GetById(id int) (domain.ConsultaDto, error)
	GetByRg(rg string) ([]domain.ConsultaDto, error)
	Update(id int, c domain.Consulta) (domain.ConsultaDto, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Post(c domain.Consulta) (domain.ConsultaDto, error) {
	return s.r.Post(c)
}

func (s *service) GetAll() []domain.ConsultaDto {
	return s.r.GetAll()
}

func (s *service) GetById(id int) (domain.ConsultaDto, error) {
	return s.r.GetById(id)
}

func (s *service) GetByRg(rg string) ([]domain.ConsultaDto, error) {
	return s.r.GetByRg(rg)
}

func (s *service) Update(id int, c domain.Consulta) (domain.ConsultaDto, error) {
	consulta, err := s.r.GetById(id)

	if err != nil {
		return domain.ConsultaDto{}, err
	}

	if c.Descricao == "" {
		c.Descricao = consulta.Descricao
	}

	if c.DataHora == "" {
		c.DataHora = consulta.DataHora
	}

	if c.MatriculaDentista == "" {
		c.MatriculaDentista = consulta.MatriculaDentista
	}

	if c.RegistroPaciente == "" {
		c.RegistroPaciente = consulta.RegistroPaciente
	}

	return s.r.Update(id, c)
}

func (s *service) Delete(id int) error {
	return s.r.Delete(id)
}