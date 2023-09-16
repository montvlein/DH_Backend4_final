package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/montvlein/DH_Backend4_final/internal/appointments"
)

type AppointmentsGetter interface {
	GetAppointmentByID(id uint) (appointments.Appointment, error)
	GetAppointmentByIDWithEntities(id uint) (appointments.AppointmentGetWithAllEntities, error)
	GetAppointmentsByPatientDNI(dni string) ([]appointments.AppointmentGetWithAllEntities, error)
}
type AppointmentCreator interface {
	CreateAppointment(appointment appointments.Appointment) (appointments.Appointment, error)
	CreateAppointmentByDNIAndLicense(appointment appointments.AppointmentCreateWithDNIAndLicense) (appointments.Appointment, error)
	ModifyAppointmentByID(id uint, appointment appointments.Appointment) (appointments.Appointment, error)
	DeleteAppointmentByID(id uint) error
}

type AppointmentsHandler struct {
	appointmentsGetter  AppointmentsGetter
	appointmentsCreator AppointmentCreator
}

func NewAppointmentsHandler(getter AppointmentsGetter, creator AppointmentCreator) *AppointmentsHandler {
	return &AppointmentsHandler{
		appointmentsGetter:  getter,
		appointmentsCreator: creator,
	}
}

// GET: traer turno por ID.
// @Summary      GET: traer turno por ID.
// @Description  Obtiene un turno por ID desde el repositorio
// @Tags         turnos
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} appointments.Appointment
// @Router       /appointments/{id} [get]
func (ph *AppointmentsHandler) GetAppointmentByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	appointment, err := ph.appointmentsGetter.GetAppointmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}
	ctx.JSON(http.StatusOK, appointment)
}

// PUT: actualizar turno.
// @Summary      PUT: actualizar turno.
// @Description  Actualiza un turno por ID desde el repositorio
// @Tags         turnos
// @Produce      json
// @Param        id path uint true "Appointment ID"
// @Param        requestBody body appointments.Appointment true "Appointment object"
// @Param PRIVATE-KEY header string true "PRIVATE-KEY"
// @Param PUBLIC-KEY header string true "PUBLIC-KEY"
// @Success      200 {object} appointments.Appointment
// @Router       /appointments/{id} [put]
func (ph *AppointmentsHandler) PutAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	appointmentRequest := appointments.Appointment{}
	err = ctx.BindJSON(&appointmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = ph.appointmentsGetter.GetAppointmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	appointment, err := ph.appointmentsCreator.ModifyAppointmentByID(uint(id), appointmentRequest)
	if err != nil {
		fmt.Println("Error when updating appointment:", err)
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, appointment)
}

// POST: agregar turno.
// @Summary      POST: agregar turno.
// @Description  Crea un nuevo turno en el sistema
// @Tags         turnos
// @Produce      json
// @Param        requestBody body appointments.Appointment true "Appointment object"
// @Param PRIVATE-KEY header string true "PRIVATE-KEY"
// @Param PUBLIC-KEY header string true "PUBLIC-KEY"
// @Success      201 {object} appointments.Appointment
// @Router       /appointments/ [post]
func (ph *AppointmentsHandler) CreateAppointment(ctx *gin.Context) {
	appointmentRequest := appointments.Appointment{}
	err := ctx.BindJSON(&appointmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(appointmentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAppointment, err := ph.appointmentsCreator.CreateAppointment(appointmentRequest)
	if err != nil {
		fmt.Println("Error when creating appointment:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, createdAppointment)
}

// DELETE: eliminar turno.
// @Summary      DELETE: eliminar turno.
// @Description  Elimina un turno por ID desde el repositorio
// @Tags         turnos
// @Produce      json
// @Param        id path uint true "Appointment ID"
// @Param PRIVATE-KEY header string true "PRIVATE-KEY"
// @Param PUBLIC-KEY header string true "PUBLIC-KEY"
// @Success      200 {string} message "Appointment deleted successfully"
// @Router       /appointments/{id} [delete]
func (ph *AppointmentsHandler) DeleteAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.appointmentsGetter.GetAppointmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	err = ph.appointmentsCreator.DeleteAppointmentByID(uint(id))
	if err != nil {
		fmt.Println("Error when deleting appointment:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

// PATCH: actualizar un turno por alguno de sus campos.
// @Summary      PATCH: actualizar un turno por alguno de sus campos.
// @Description  Actualiza un turno parcialmente por ID desde el repositorio
// @Tags         turnos
// @Produce      json
// @Param        id path uint true "Appointment ID"
// @Param        requestBody body appointments.AppointmentPatch true "Appointment patch object"
// @Param PRIVATE-KEY header string true "PRIVATE-KEY"
// @Param PUBLIC-KEY header string true "PUBLIC-KEY"
// @Success      200 {object} appointments.Appointment
// @Router       /appointments/{id} [patch]
func (ph *AppointmentsHandler) PatchAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	existingAppointment, err := ph.appointmentsGetter.GetAppointmentByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	var appointmentPatch appointments.AppointmentPatch
	if err := ctx.BindJSON(&appointmentPatch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if appointmentPatch.Description != "" {
		existingAppointment.Description = appointmentPatch.Description
	}
	if appointmentPatch.Date != "" {
		existingAppointment.Date = appointmentPatch.Date
	}

	if appointmentPatch.Hour != "" {
		existingAppointment.Hour = appointmentPatch.Hour
	}

	updatedAppointment, err := ph.appointmentsCreator.ModifyAppointmentByID(uint(id), existingAppointment)
	if err != nil {
		fmt.Println("Error when updating appointment:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, updatedAppointment)
}

// GET: traer turno por ID pero con el detalle de las entidades
// @Summary      GET: traer turno por ID pero con el detalle de las entidades
// @Description  Obtiene un turno por ID con el detalle de las entidades de paciente y dentista desde el repositorio
// @Tags         turnos
// @Produce      json
// @Param        id path int true "Appointment ID"
// @Success      200 {object} appointments.AppointmentGetWithAllEntities
// @Router       /appointments/{id}/details [get]
func (ph *AppointmentsHandler) GetAppointmentByIDWithEntities(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	appointment, err := ph.appointmentsGetter.GetAppointmentByIDWithEntities(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

// POST: agregar turno por DNI del paciente y matrícula del dentista.
// @Summary      POST: agregar turno por DNI del paciente y matrícula del dentista.
// @Description  Crea un nuevo turno asociado a un paciente identificado por su DNI y a un dentista identificado por su matrícula
// @Tags         turnos
// @Produce      json
// @Param        requestBody body appointments.AppointmentCreateWithDNIAndLicense true "Appointment data"
// @Param PRIVATE-KEY header string true "PRIVATE-KEY"
// @Param PUBLIC-KEY header string true "PUBLIC-KEY"
// @Success      201 {object} appointments.Appointment
// @Router       /appointments/create-by-dni-and-license [post]
func (ph *AppointmentsHandler) CreateAppointmentByDNIAndLicense(ctx *gin.Context) {
	var appointment appointments.AppointmentCreateWithDNIAndLicense

	if err := ctx.BindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAppointment, err := ph.appointmentsCreator.CreateAppointmentByDNIAndLicense(appointment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAppointment)
}

// GET: traer turno por DNI del paciente. Debe traer el detalle del turno (Fecha-Hora, descripción, Paciente y Dentista) y el dni deberá ser recibido por QueryParams.
// @Summary GET: traer turno por DNI del paciente. Debe traer el detalle del turno (Fecha-Hora, descripción, Paciente y Dentista) y el dni deberá ser recibido por QueryParams.
// @Description Obtiene los turnos de un paciente por su DNI con el detalle de las entidades de paciente y dentista desde el repositorio
// @Tags turnos
// @Produce json
// @Param dni query string true "Patient's DNI"
// @Success 200 {array} appointments.AppointmentGetWithAllEntities
// @Router /appointments/patient [get]
func (ph *AppointmentsHandler) GetAppointmentsByPatientDNI(ctx *gin.Context) {
	patientDNI := ctx.Query("dni")
	if patientDNI == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dni parameter is required"})
		return
	}

	appointments, err := ph.appointmentsGetter.GetAppointmentsByPatientDNI(patientDNI)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}
