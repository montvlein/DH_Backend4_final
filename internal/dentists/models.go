package dentists

type Dentist struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Lastname string `json:"lastname" validate:"required"`
	License  string `json:"license" validate:"required"`
}

type DentistPatch struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	License  string `json:"license"`
}
