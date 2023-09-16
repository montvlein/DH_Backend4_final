package appointments

type Appointment struct {
	ID          uint   `json:"id"`
	PatientID   uint   `json:"patient_id" validate:"required"`
	DentistID   uint   `json:"dentist_id" validate:"required"`
	Description string `json:"description"`
	Date        string `json:"date" validate:"required"`
	Hour        string `json:"hour" validate:"required"`
}

type AppointmentPatch struct {
	Description string `json:"description"`
	Date        string `json:"date"`
	Hour        string `json:"hour"`
}

type AppointmentGetWithAllEntities struct {
	ID          uint        `json:"id"`
	Patient     PatientInfo `json:"patient"`
	Dentist     DentistInfo `json:"dentist"`
	Description string      `json:"description"`
	Date        string      `json:"date"`
	Hour        string      `json:"hour"`
}

type PatientInfo struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Lastname      string `json:"lastname"`
	Address       string `json:"address"`
	DNI           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}

type DentistInfo struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	License  string `json:"license"`
}

type AppointmentCreateWithDNIAndLicense struct {
	ID             uint   `json:"id"`
	PatientDNI     string `json:"patient_dni" validate:"required"`
	DentistLicense string `json:"dentist_license" validate:"required"`
	Description    string `json:"description"`
	Date           string `json:"date" validate:"required"`
	Hour           string `json:"hour" validate:"required"`
}
