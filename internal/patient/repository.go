package patient

type RepositoryPatient interface {
	//PutPatient(patient Patient) (Patient, error)
	GetPatientByID(id int) (Patient, error)
	ModifyPatientByID(id int, patient Patient) (Patient, error)
	//ModifyPatientParcialByID(id int, patient Patient) (Patient, error)
	//DeletePatientByID(id int)(Patient, error)
}
