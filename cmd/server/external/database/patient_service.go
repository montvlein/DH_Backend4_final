package database

import (
	"fmt"

	"github.com/montvlein/DH_Backend4_final/internal/patients"
)

func (s *SqlStore) GetPatientById(id int) (patients.Patient, error) {
	var patientReturn patients.Patient

	query := fmt.Sprintf("SELECT * FROM patients WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&patientReturn.ID, &patientReturn.Name, &patientReturn.Lastname, &patientReturn.Address, &patientReturn.DNI, &patientReturn.DischargeDate)
	if err != nil {
		return patients.Patient{}, err
	}
	return patientReturn, nil
}

func (s *SqlStore) GetPatientList() ([]patients.Patient, error) {
	var patientsReturn []patients.Patient

	query := fmt.Sprintf("SELECT * FROM patients;")
	rows, err := s.DB.Query(query)
	if err != nil {
		return []patients.Patient{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var patient patients.Patient
		err := rows.Scan(&patient.ID, &patient.Name, &patient.Lastname, &patient.Address, &patient.DNI, &patient.DischargeDate)
		if err != nil {
			return []patients.Patient{}, err
		}
		patientsReturn = append(patientsReturn, patient)
	}

	if err != nil {
		return []patients.Patient{}, err
	}
	return patientsReturn, nil
}

func (s *SqlStore) ModifyPatientById(id int, patient patients.Patient) (patients.Patient, error) {
	query := fmt.Sprintf("UPDATE patients SET name = '%s', lastname = '%v', address = '%s', dni = '%s', discharge_date = '%s' WHERE id = %v;", patient.Name, patient.Lastname,
		patient.Address, patient.DNI, patient.DischargeDate, id)

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return patients.Patient{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return patients.Patient{}, err
	}

	patient.ID = uint(id)
	return patient, nil
}

func (s *SqlStore) CreatePatient(patient patients.Patient) (patients.Patient, error) {
	query := fmt.Sprintf("INSERT INTO patients (name, lastname, address, dni, discharge_date) VALUES ('%s', '%s', '%s', '%s', '%s') RETURNING id;", patient.Name, patient.Lastname, patient.Address, patient.DNI, patient.DischargeDate)
	var id int
	err := s.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return patients.Patient{}, err
	}
	patient.ID = uint(id)
	return patient, nil
}

func (s *SqlStore) DeletePatientById(id int) error {
	query := fmt.Sprintf("DELETE FROM patients WHERE id = %d;", id)
	_, err := s.DB.Exec(query)
	return err
}
