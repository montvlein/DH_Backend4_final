package dentist

// Repository is an interface that we used to indicate to some user how to implement
// a repository for dentists.
type RepositoryDentist interface {
	
	//PutDentist(dentist Dentist) (Dentist, error)
	GetDentistByID(id int) (Dentist, error)
	ModifyDentistByID(id int, dentist Dentist) (Dentist, error)
	//ModifyDentistParcialByID(id int, dentist Dentist) (Dentist, error)
	//DeleteDentistByID(id int) (Dentist, error)
}