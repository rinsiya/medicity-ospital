package handler

import (
	//"medicity/internal/dto"
	"medicity/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type patientHandler struct {
	patientService service.PatientService
}

func NewPatientHandler(patientService service.PatientService) *patientHandler {
	return &patientHandler{
		patientService: patientService,
	}
}


func (h *patientHandler) PatientSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "patientSignup.html", nil)
}



