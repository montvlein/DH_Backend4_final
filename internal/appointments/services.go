package appointments

type Repository interface {
	GetAppointmentByID(id uint) (Appointment, error)
	GetAppointmentByIDWithEntities(id uint) (AppointmentGetWithAllEntities, error)
	CreateAppointment(appointment Appointment) (Appointment, error)
	CreateAppointmentByDNIAndLicense(appointment AppointmentCreateWithDNIAndLicense) (Appointment, error)
	ModifyAppointmentByID(id uint, appointment Appointment) (Appointment, error)
	DeleteAppointmentByID(id uint) error
	GetAppointmentsByPatientDNI(patientID string) ([]AppointmentGetWithAllEntities, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAppointmentByID(id uint) (Appointment, error) {
	return s.repository.GetAppointmentByID(id)
}

func (s *Service) GetAppointmentByIDWithEntities(id uint) (AppointmentGetWithAllEntities, error) {
	return s.repository.GetAppointmentByIDWithEntities(id)
}

func (s *Service) ModifyAppointmentByID(id uint, appointment Appointment) (Appointment, error) {
	return s.repository.ModifyAppointmentByID(id, appointment)
}

func (s *Service) CreateAppointment(appointment Appointment) (Appointment, error) {
	return s.repository.CreateAppointment(appointment)
}

func (s *Service) CreateAppointmentByDNIAndLicense(appointment AppointmentCreateWithDNIAndLicense) (Appointment, error) {
	return s.repository.CreateAppointmentByDNIAndLicense(appointment)
}

func (s *Service) DeleteAppointmentByID(id uint) error {
	return s.repository.DeleteAppointmentByID(id)
}

func (s *Service) GetAppointmentsByPatientDNI(patientID string) ([]AppointmentGetWithAllEntities, error) {
	return s.repository.GetAppointmentsByPatientDNI(patientID)
}
