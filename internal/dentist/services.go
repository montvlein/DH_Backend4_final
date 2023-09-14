package dentist

// Service provides all functionalities related to dentists.
type ServiceDentist struct {
	repositoryDentist RepositoryDentist
}

func NewServiceDentist(repositoryDentist RepositoryDentist) *ServiceDentist {
	return &ServiceDentist{repositoryDentist: repositoryDentist}
}

//func (s *ServiceDentist) PutDentist( dentist Dentist) (Dentist, error) {
//	return s.repository.PutDentist(id, dentist)
//}

func (s *ServiceDentist) GetDentistByID(id int) (Dentist, error) {
	return s.repositoryDentist.GetDentistByID(id)
}

func (s *ServiceDentist) ModifyDentistByID(id int, dentist Dentist) (Dentist, error) {
	return s.repositoryDentist.ModifyDentistByID(id, dentist)
}

//func (s *ServiceDentist) ModifyDentistParcialByID(id int, dentist Dentist) (Dentist, error) {
//	return s.repositoryDentist.ModifyDentistParcialByID(id, dentist)
//}

//func (s *ServiceDentist) DeleteDentistByID(id int, dentist Dentist) (Dentist, error) {
//	return s.repositoryDentist.DeleteDentistByID(id)
//}