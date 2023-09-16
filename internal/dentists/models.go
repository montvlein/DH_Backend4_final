package dentists

// Dentist represents a dentist entity.
type Dentist struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	License  string `json:"license" validate:"required"`
}

// Dentist represents a dentist entity.
type DentistPatch struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	License  string `json:"license"`
}
