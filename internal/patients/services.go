package patients

type Repository interface {
	GetPatientById(id int) (Patient, error)
	CreatePatient(patient Patient) (Patient, error)
	ModifyPatientById(id int, patient Patient) (Patient, error)
	DeletePatientById(id int) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetPatientByID(id int) (Patient, error) {
	return s.repository.GetPatientById(id)
}

func (s *Service) ModifyPatientByID(id int, patient Patient) (Patient, error) {
	return s.repository.ModifyPatientById(id, patient)
}

func (s *Service) CreatePatient(patient Patient) (Patient, error) {
	return s.repository.CreatePatient(patient)
}

func (s *Service) DeletePatientByID(id int) error {
	return s.repository.DeletePatientById(id)
}
