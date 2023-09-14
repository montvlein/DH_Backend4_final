package patients

type Patient struct {
	ID            uint   `json:"id"`
	Name          string `json:"name" validate:"required"`
	Lastname      string `json:"lastname" validate:"required"`
	Address       string `json:"address" validate:"required"`
	DNI           string `json:"dni" validate:"required"`
	DischargeDate string `json:"discharge_date" validate:"required"`
}

type PatientPatch struct {
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	Address       string `json:"address"`
	DNI           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}
