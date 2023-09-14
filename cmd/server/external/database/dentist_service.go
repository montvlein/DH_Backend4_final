package database

import (
	"database/sql"
	"fmt"

	"github.com/montvlein/DH_Backend4_final/internal/dentists"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetById(id int) (dentists.Dentist, error) {
	var dentistReturn dentists.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.Lastname, &dentistReturn.License)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) ModifyById(id int, dentist dentists.Dentist) (dentists.Dentist, error) {
	query := fmt.Sprintf("UPDATE dentists SET name = '%s', lastname = '%v', license = '%s' WHERE id = %v;", dentist.Name, dentist.Lastname,
		dentist.License, id)

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return dentists.Dentist{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return dentists.Dentist{}, err
	}

	dentist.ID = uint(id)
	return dentist, nil
}

func (s *SqlStore) Create(dentist dentists.Dentist) (dentists.Dentist, error) {
	query := fmt.Sprintf("INSERT INTO dentists (name, lastname, license) VALUES ('%s', '%s', '%s') RETURNING id;", dentist.Name, dentist.Lastname, dentist.License)
	var id int
	err := s.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return dentists.Dentist{}, err
	}
	dentist.ID = uint(id)
	return dentist, nil
}

func (s *SqlStore) DeleteById(id int) error {
	query := fmt.Sprintf("DELETE FROM dentists WHERE id = %d;", id)
	_, err := s.DB.Exec(query)
	return err
}
