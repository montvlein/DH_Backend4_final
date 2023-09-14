package memory

import (
	"DH_Backend4_final/internal/dentist"
	"encoding/json"
	"errors"
	"os"
)

// Service provides functionalities related to memory dentists storage
type ServiceDentist struct {
	dentistList []dentist.Dentist
}

func NewServiceDentist(path string) (*ServiceDentist, error) {
	byteFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var dentists []dentist.Dentist
	err = json.Unmarshal(byteFile, &dentists)
	if err != nil {
		return nil, err
	}
	return &ServiceDentist{dentistList: dentists}, nil
}

func (s *ServiceDentist) GetDentistByID(id int) (dentist.Dentist, error) {
	for _, value := range s.dentistList {
		if value.ID == id {
			return value, nil
		}
	}
	return dentist.Dentist{}, errors.New("not found")
}

func (s *ServiceDentist) ModifyDentistByID(id int, dentistmod dentist.Dentist) (dentist.Dentist, error) {
	for i, value := range s.dentistList {
		if value.ID == id {
			s.dentistList = append(s.dentistList[:i], s.dentistList[i+1:]...)
			s.dentistList = append(s.dentistList, dentistmod)
			return dentistmod, nil
		}
	}
	return dentist.Dentist{}, errors.New("not found")
}
