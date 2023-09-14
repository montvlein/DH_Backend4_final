package dentists

type Repository interface {
	GetDentistById(id int) (Dentist, error)
	CreateDentist(dentist Dentist) (Dentist, error)
	ModifyDentistById(id int, dentist Dentist) (Dentist, error)
	DeleteDentistById(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetDentistByID(id int) (Dentist, error) {
	return s.repository.GetDentistById(id)
}

func (s *Service) ModifyDentistByID(id int, dentist Dentist) (Dentist, error) {
	return s.repository.ModifyDentistById(id, dentist)
}

func (s *Service) CreateDentist(dentist Dentist) (Dentist, error) {
	return s.repository.CreateDentist(dentist)
}

func (s *Service) DeleteDentistByID(id int) error {
	return s.repository.DeleteDentistById(id)
}
