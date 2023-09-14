package handler

import (
	"DH_Backend4_final/internal/dentist"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type DentistsGetter interface {
	GetDentistByID(id int) (dentist.Dentist, error)
}

type DentistModificator interface {
	ModifyDentistByID(id int, dentist dentist.Dentist) (dentist.Dentist, error)
}

type DentistsHandler struct {
	dentistsGetter  DentistsGetter
	DentistModificator DentistModificator
}

func NewDentistsHandler(getter DentistsGetter, modificator DentistModificator) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		DentistModificator: modificator,
	}
}

// GetDentistByID godoc
// @Summary      Gets a dentist by id
// @Description  Gets a dentist by id from the repository
// @Tags         dentists
// @Produce      json
// @Param        id path string true "ID"
// @Success      200 {object} dentist.Dentist
// @Router       /dentists/{id} [get]
func (dh *DentistsHandler) GetDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	dentist, err := dh.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "dentist not found"})
		return
	}
	ctx.JSON(http.StatusOK, dentist)
}

func (dh *DentistsHandler) PutDentist(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	dentistRequest := dentist.Dentist{}
	err = ctx.BindJSON(&dentistRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dentist, err := dh.DentistModificator.ModifyDentistByID(id, dentistRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, dentist)
}
