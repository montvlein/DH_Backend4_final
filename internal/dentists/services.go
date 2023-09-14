package dentists

type Repository interface {
	GetById(id int) (Dentist, error)
	Create(dentist Dentist) (Dentist, error)
	ModifyById(id int, dentist Dentist) (Dentist, error)
	DeleteById(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Dentist, error) {
	return s.repository.GetById(id)
}

func (s *Service) ModifyByID(id int, dentist Dentist) (Dentist, error) {
	return s.repository.ModifyById(id, dentist)
}

func (s *Service) CreateDentist(dentist Dentist) (Dentist, error) {
	return s.repository.Create(dentist)
}

func (s *Service) DeleteDentistByID(id int) error {
	return s.repository.DeleteById(id)
}
