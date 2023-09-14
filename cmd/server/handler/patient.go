package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/montvlein/DH_Backend4_final/internal/patients"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PatientsGetter interface {
	GetPatientByID(id int) (patients.Patient, error)
}
type PatientCreator interface {
	CreatePatient(patient patients.Patient) (patients.Patient, error)
	ModifyPatientByID(id int, patient patients.Patient) (patients.Patient, error)
	DeletePatientByID(id int) error
}

type PatientsHandler struct {
	patientsGetter  PatientsGetter
	patientsCreator PatientCreator
}

func NewPatientsHandler(getter PatientsGetter, creator PatientCreator) *PatientsHandler {
	return &PatientsHandler{
		patientsGetter:  getter,
		patientsCreator: creator,
	}
}

// GetPatientByID godoc
// @Summary      Gets a patient by id
// @Description  Gets a patient by id from the repository
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} patients.Patient
// @Router       /patients/{id} [get]
func (ph *PatientsHandler) GetPatientByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	patient, err := ph.patientsGetter.GetPatientByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

// PutPatient godoc
// @Summary      Updates a patient by ID
// @Description  Updates a patient by ID from the repository
// @Tags         patients
// @Produce      json
// @Param        id path int true "Patient ID"
// @Param        requestBody body patients.Patient true "Patient object"
// @Success      200 {object} patients.Patient
// @Router       /patients/{id} [put]
func (ph *PatientsHandler) PutPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	patientRequest := patients.Patient{}
	err = ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = ph.patientsGetter.GetPatientByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	patient, err := ph.patientsCreator.ModifyPatientByID(id, patientRequest)
	if err != nil {
		fmt.Println("Error when updating patient:", err)
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, patient)
}

// CreatePatient godoc
// @Summary      Creates a new patient
// @Description  Creates a new patient in the repository
// @Tags         patients
// @Produce      json
// @Param        requestBody body patients.Patient true "Patient object"
// @Success      201 {object} patients.Patient
// @Router       /patients/ [post]
func (ph *PatientsHandler) CreatePatient(ctx *gin.Context) {
	patientRequest := patients.Patient{}
	err := ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPatient, err := ph.patientsCreator.CreatePatient(patientRequest)
	if err != nil {
		fmt.Println("Error when creating patient:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, createdPatient)
}

// DeletePatient godoc
// @Summary      Deletes a patient by ID
// @Description  Deletes a patient by ID from the repository
// @Tags         patients
// @Produce      json
// @Param        id path int true "Patient ID"
// @Success      200 {string} message "Patient deleted successfully"
// @Router       /patients/{id} [delete]
func (ph *PatientsHandler) DeletePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.patientsGetter.GetPatientByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	err = ph.patientsCreator.DeletePatientByID(id)
	if err != nil {
		fmt.Println("Error when deleting patient:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

// PatchPatient godoc
// @Summary      Updates a patient partially by ID
// @Description  Updates a patient partially by ID from the repository
// @Tags         patients
// @Produce      json
// @Param        id path int true "Patient ID"
// @Param        requestBody body patients.PatientPatch true "Patient patch object"
// @Success      200 {object} patients.Patient
// @Router       /patients/{id} [patch]
func (ph *PatientsHandler) PatchPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	existingPatient, err := ph.patientsGetter.GetPatientByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}

	var patientPatch patients.PatientPatch
	if err := ctx.BindJSON(&patientPatch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if patientPatch.Name != "" {
		existingPatient.Name = patientPatch.Name
	}
	if patientPatch.Lastname != "" {
		existingPatient.Lastname = patientPatch.Lastname
	}
	if patientPatch.Address != "" {
		existingPatient.Address = patientPatch.Address
	}
	if patientPatch.DNI != "" {
		existingPatient.DNI = patientPatch.DNI
	}
	if patientPatch.DischargeDate != "" {
		existingPatient.DischargeDate = patientPatch.DischargeDate
	}

	updatedPatient, err := ph.patientsCreator.ModifyPatientByID(id, existingPatient)
	if err != nil {
		fmt.Println("Error when updating patient:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, updatedPatient)
}
