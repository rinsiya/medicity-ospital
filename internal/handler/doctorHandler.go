package handler

import (
	//"medicity/internal/dto"
	"medicity/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type doctorHandler struct {
	doctorService service.DoctorService
}

func NewDoctorHandler(doctorService service.DoctorService) *doctorHandler {
	return &doctorHandler{
		doctorService: doctorService,
	}
}


func (h *doctorHandler) DoctorSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "doctorSignup.html", nil)
}



