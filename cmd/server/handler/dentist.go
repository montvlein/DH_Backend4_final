package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/montvlein/DH_Backend4_final/internal/dentists"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DentistsGetter interface {
	GetDentistByID(id int) (dentists.Dentist, error)
}
type DentistCreator interface {
	CreateDentist(dentist dentists.Dentist) (dentists.Dentist, error)
	ModifyDentistByID(id int, dentist dentists.Dentist) (dentists.Dentist, error)
	DeleteDentistByID(id int) error
}

type DentistsHandler struct {
	dentistsGetter  DentistsGetter
	dentistsCreator DentistCreator
}

func NewDentistsHandler(getter DentistsGetter, creator DentistCreator) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		dentistsCreator: creator,
	}
}

// GetDentistByID
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} dentists.Dentist
// @Router       /dentists/{id} [get]
func (ph *DentistsHandler) GetDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	dentist, err := ph.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

// PutDentist
// @Summary      Updates a dentist by ID
// @Description  Updates a dentist by ID from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Param        requestBody body Dentist true "Dentist object"
// @Success      200 {object} dentists.Dentist
// @Router       /dentists/{id} [put]
func (ph *DentistsHandler) PutDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	dentistRequest := dentists.Dentist{}
	err = ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = ph.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}

	dentist, err := ph.dentistsCreator.ModifyDentistByID(id, dentistRequest)
	if err != nil {
		fmt.Println("Error when updating dentist:", err)
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, dentist)
}

// CreateDentist
// @Summary      Creates a new dentist
// @Description  Creates a new dentist in the repository
// @Tags         dentists
// @Produce      json
// @Param        requestBody body Dentist true "Dentist object"
// @Success      201 {object} dentists.Dentist
// @Router       /dentists/ [post]
func (ph *DentistsHandler) CreateDentist(ctx *gin.Context) {
	dentistRequest := dentists.Dentist{}
	err := ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDentist, err := ph.dentistsCreator.CreateDentist(dentistRequest)
	if err != nil {
		fmt.Println("Error when creating dentist:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusCreated, createdDentist)
}

// DeleteDentist
// @Summary      Deletes a dentist by ID
// @Description  Deletes a dentist by ID from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Success      200 {string} message "Dentist deleted successfully"
// @Router       /dentists/{id} [delete]
func (ph *DentistsHandler) DeleteDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	_, err = ph.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}

	err = ph.dentistsCreator.DeleteDentistByID(id)
	if err != nil {
		fmt.Println("Error when deleting dentist:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Dentist deleted successfully"})
}

// PatchDentist
// @Summary      Updates a dentist partially by ID
// @Description  Updates a dentist partially by ID from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path int true "Dentist ID"
// @Param        requestBody body DentistPatch true "Dentist patch object"
// @Success      200 {object} dentists.Dentist
// @Router       /dentists/{id} [patch]
func (ph *DentistsHandler) PatchDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	existingDentist, err := ph.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}

	var dentistPatch dentists.DentistPatch
	if err := ctx.BindJSON(&dentistPatch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dentistPatch.Name != "" {
		existingDentist.Name = dentistPatch.Name
	}
	if dentistPatch.Lastname != "" {
		existingDentist.Lastname = dentistPatch.Lastname
	}
	if dentistPatch.License != "" {
		existingDentist.License = dentistPatch.License
	}

	updatedDentist, err := ph.dentistsCreator.ModifyDentistByID(id, existingDentist)
	if err != nil {
		fmt.Println("Error when updating dentist:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, updatedDentist)
}
