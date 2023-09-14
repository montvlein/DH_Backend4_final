package memory

import (
	"DH_Backend4_final/internal/patient"
	"encoding/json"
	"errors"
	"os"
)

// Service provides functionalities related to memory dentists storage
type ServicePatient struct {
	patientList []patient.Patient
}

func NewServicePatient(path string) (*ServicePatient, error) {
	byteFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var patients []patient.Patient
	err = json.Unmarshal(byteFile, &patients)
	if err != nil {
		return nil, err
	}
	return &ServicePatient{patientList: patients}, nil
}

func (s *ServicePatient) GetPatientByID(id int) (patient.Patient, error) {
	for _, value := range s.patientList {
		if value.ID == id {
			return value, nil
		}
	}
	return patient.Patient{}, errors.New("not found")
}

func (s *ServicePatient) ModifyPatientByID(id int, patientmod patient.Patient) (patient.Patient, error) {
	for i, value := range s.patientList {
		if value.ID == id {
			s.patientList = append(s.patientList[:i], s.patientList[i+1:]...)
			s.patientList = append(s.patientList, patientmod)
			return patientmod, nil
		}
	}
	return patient.Patient{}, errors.New("not found")
}
