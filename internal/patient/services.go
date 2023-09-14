package patient

// Service provides all functionalities related to patients.
type ServicePatient struct {
	repositoryPatient RepositoryPatient
}

func NewServicePatient(repositoryPatient RepositoryPatient) *ServicePatient {
	return &ServicePatient{repositoryPatient: repositoryPatient}
}

//func (s *Service) PutPatient(patient Patient) (Patient, error) {
//	return s.repositoryPatient.PutPatient(id, Patient)
//}

func (s *ServicePatient) GetPatientByID(id int) (Patient, error) {
	return s.repositoryPatient.GetPatientByID(id)
}

func (s *ServicePatient) ModifyPatientByID(id int, patient Patient) (Patient, error) {
	return s.repositoryPatient.ModifyPatientByID(id, patient)
}

//func (s *ServicePatient) ModifyPatientParcialByID(id int, patient Patient) (Patient, error) {
//	return s.repositoryPatient.ModifyParcialByID(id, patient)
//}

//func (s *ServicePatient) DeletePatientByID(id int, patient Patient) (Patient, error) {
//	return s.repositoryPatient.DeleteByID(id)
//}
