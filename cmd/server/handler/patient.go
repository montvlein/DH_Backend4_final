package handler

import (
	"DH_Backend4_final/internal/patient"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientsGetter interface {
	GetPatientByID(id int) (patient.Patient, error)
}

type PatientModificator interface {
	ModifyPatientByID(id int, patient patient.Patient) (patient.Patient, error)
}

type PatientsHandler struct {
	patientsGetter     PatientsGetter
	PatientModificator PatientModificator
}

func NewPatientsHandler(getter PatientsGetter, modificator PatientModificator) *PatientsHandler {
	return &PatientsHandler{
		patientsGetter:     getter,
		PatientModificator: modificator,
	}
}

// GetPatientByID godoc
// @Summary      Gets a patient by id
// @Description  Gets a patient by id from the repository
// @Tags         patients
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} patient.Patient
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

func (ph *PatientsHandler) PutPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	patientRequest := patient.Patient{}
	err = ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := ph.PatientModificator.ModifyPatientByID(id, patientRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, patient)
}
