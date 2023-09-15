package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/montvlein/DH_Backend4_final/internal/appointments"
)

func (s *SqlStore) GetAppointmentByID(id uint) (appointments.Appointment, error) {
	var appointmentReturn appointments.Appointment

	query := fmt.Sprintf("SELECT * FROM appointments WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&appointmentReturn.ID, &appointmentReturn.PatientID, &appointmentReturn.DentistID, &appointmentReturn.Description, &appointmentReturn.Date, &appointmentReturn.Hour)
	if err != nil {
		return appointments.Appointment{}, err
	}
	return appointmentReturn, nil
}

func (s *SqlStore) ModifyAppointmentByID(id uint, appointment appointments.Appointment) (appointments.Appointment, error) {
	query := fmt.Sprintf("UPDATE appointments SET description = '%s', date = '%v', hour = '%s' WHERE id = %v;", appointment.Description, appointment.Date, appointment.Hour, id)

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appointments.Appointment{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return appointments.Appointment{}, err
	}

	appointment.ID = id
	return appointment, nil
}

func (s *SqlStore) CreateAppointment(appointment appointments.Appointment) (appointments.Appointment, error) {
	query := fmt.Sprintf("INSERT INTO appointments (patient_id, dentist_id, description, date, hour) VALUES (%d, %d, '%s', '%v', '%s') RETURNING id;", appointment.PatientID, appointment.DentistID, appointment.Description, appointment.Date, appointment.Hour)
	var id uint
	err := s.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return appointments.Appointment{}, err
	}
	appointment.ID = id
	return appointment, nil
}

func (s *SqlStore) DeleteAppointmentByID(id uint) error {
	query := fmt.Sprintf("DELETE FROM appointments WHERE id = %d;", id)
	_, err := s.DB.Exec(query)
	return err
}

func (s *SqlStore) GetAppointmentByIDWithEntities(id uint) (appointments.AppointmentGetWithAllEntities, error) {
	var appointmentTemp appointments.Appointment
	var appointmentReturn appointments.AppointmentGetWithAllEntities

	appointmentQuery := fmt.Sprintf("SELECT id, description, date, hour, patient_id, dentist_id FROM appointments WHERE id = %d;", id)
	appointmentRow := s.DB.QueryRow(appointmentQuery)
	err := appointmentRow.Scan(
		&appointmentTemp.ID, &appointmentTemp.Description, &appointmentTemp.Date, &appointmentTemp.Hour,
		&appointmentTemp.PatientID, &appointmentTemp.DentistID,
	)
	if err != nil {
		return appointments.AppointmentGetWithAllEntities{}, err
	}
	appointmentReturn.ID = appointmentTemp.ID
	appointmentReturn.Date = appointmentTemp.Date
	appointmentReturn.Description = appointmentTemp.Description
	appointmentReturn.Hour = appointmentTemp.Hour

	patientQuery := fmt.Sprintf("SELECT id, name, lastname, address, dni, discharge_date FROM patients WHERE id = %d;", appointmentTemp.PatientID)
	patientRow := s.DB.QueryRow(patientQuery)
	patientInfo := appointments.PatientInfo{}
	err = patientRow.Scan(
		&patientInfo.ID, &patientInfo.Name, &patientInfo.Lastname, &patientInfo.Address, &patientInfo.DNI, &patientInfo.DischargeDate,
	)
	if err != nil {
		return appointments.AppointmentGetWithAllEntities{}, err
	}
	appointmentReturn.Patient = patientInfo

	dentistQuery := fmt.Sprintf("SELECT id, name, lastname, license FROM dentists WHERE id = %d;", appointmentTemp.DentistID)
	dentistRow := s.DB.QueryRow(dentistQuery)
	dentistInfo := appointments.DentistInfo{}
	err = dentistRow.Scan(
		&dentistInfo.ID, &dentistInfo.Name, &dentistInfo.Lastname, &dentistInfo.License,
	)
	if err != nil {
		return appointments.AppointmentGetWithAllEntities{}, err
	}
	appointmentReturn.Dentist = dentistInfo

	return appointmentReturn, nil
}

func (s *SqlStore) CreateAppointmentByDNIAndLicense(appointment appointments.AppointmentCreateWithDNIAndLicense) (appointments.Appointment, error) {
	queryPatient := fmt.Sprintf("SELECT id FROM patients WHERE dni = '%s';", appointment.PatientDNI)
	queryDentist := fmt.Sprintf("SELECT id FROM dentists WHERE license = '%s';", appointment.DentistLicense)

	var patientID, dentistID uint
	var appointmentResult appointments.Appointment

	err := s.DB.QueryRow(queryPatient).Scan(&patientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appointments.Appointment{}, fmt.Errorf("Patient not found")
		}
		return appointments.Appointment{}, err
	}

	err = s.DB.QueryRow(queryDentist).Scan(&dentistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return appointments.Appointment{}, fmt.Errorf("Dentist not found")
		}
		return appointments.Appointment{}, err
	}

	queryInsert := fmt.Sprintf("INSERT INTO appointments (patient_id, dentist_id, description, date, hour) VALUES (%d, %d, '%s', '%s', '%s') RETURNING id;",
		patientID, dentistID, appointment.Description, appointment.Date, appointment.Hour)

	var id uint
	err = s.DB.QueryRow(queryInsert).Scan(&id)
	if err != nil {
		return appointments.Appointment{}, err
	}

	appointmentResult.ID = id
	appointmentResult.PatientID = patientID
	appointmentResult.DentistID = dentistID
	appointmentResult.Date = appointment.Date
	appointmentResult.Description = appointment.Description
	appointmentResult.Hour = appointment.Hour

	return appointmentResult, nil
}

func (s *SqlStore) GetAppointmentsByPatientDNI(patientDNI string) ([]appointments.AppointmentGetWithAllEntities, error) {
	queryPatient := fmt.Sprintf("SELECT id FROM patients WHERE dni = '%s';", patientDNI)

	var patientID uint
	var dentistID uint

	err := s.DB.QueryRow(queryPatient).Scan(&patientID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || patientID == 0 {
			return []appointments.AppointmentGetWithAllEntities{}, fmt.Errorf("Patient not found")
		}
		return []appointments.AppointmentGetWithAllEntities{}, err
	}

	queryDentist := fmt.Sprintf("SELECT dentist_id FROM appointments WHERE patient_id = %d LIMIT 1;", patientID)

	err = s.DB.QueryRow(queryDentist).Scan(&dentistID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || dentistID == 0 {
			return []appointments.AppointmentGetWithAllEntities{}, fmt.Errorf("Dentist not found")
		}
		return []appointments.AppointmentGetWithAllEntities{}, err
	}

	queryPatientInfo := fmt.Sprintf("SELECT id, name, lastname, address, dni, discharge_date FROM patients WHERE id = %d;", patientID)
	var patientInfo appointments.PatientInfo

	err = s.DB.QueryRow(queryPatientInfo).Scan(
		&patientInfo.ID, &patientInfo.Name, &patientInfo.Lastname, &patientInfo.Address, &patientInfo.DNI, &patientInfo.DischargeDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []appointments.AppointmentGetWithAllEntities{}, fmt.Errorf("Patient not found")
		}
		return []appointments.AppointmentGetWithAllEntities{}, err
	}

	queryDentistInfo := fmt.Sprintf("SELECT id, name, lastname, license FROM dentists WHERE id = %d;", dentistID)
	var dentistInfo appointments.DentistInfo

	err = s.DB.QueryRow(queryDentistInfo).Scan(
		&dentistInfo.ID, &dentistInfo.Name, &dentistInfo.Lastname, &dentistInfo.License,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []appointments.AppointmentGetWithAllEntities{}, fmt.Errorf("Dentist not found")
		}
		return []appointments.AppointmentGetWithAllEntities{}, err
	}

	queryAppointments := fmt.Sprintf("SELECT id, description, date, hour FROM appointments WHERE patient_id = %d;", patientID)
	rows, err := s.DB.Query(queryAppointments)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []appointments.AppointmentGetWithAllEntities{}, fmt.Errorf("Appointments not found")
		}
		return []appointments.AppointmentGetWithAllEntities{}, err
	}

	var appointmentsReturn []appointments.AppointmentGetWithAllEntities

	// Iterar sobre las citas y combinar la información
	for rows.Next() {
		var appointmentTemp appointments.AppointmentGetWithAllEntities
		err := rows.Scan(&appointmentTemp.ID, &appointmentTemp.Description, &appointmentTemp.Date, &appointmentTemp.Hour)
		if err != nil {
			return []appointments.AppointmentGetWithAllEntities{}, err
		}

		appointmentTemp.Patient = patientInfo
		appointmentTemp.Dentist = dentistInfo
		appointmentsReturn = append(appointmentsReturn, appointmentTemp)
	}

	return appointmentsReturn, nil
}
