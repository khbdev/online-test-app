package handler

import (
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type JobHandler struct {
	jobService *service.JobService
}

// ✅ yangi handler yaratish
func NewJobHandler(jobService *service.JobService) *JobHandler {
	return &JobHandler{jobService: jobService}
}

// ✅ SubmitTest — testni yuborish
func (h *JobHandler) SubmitTest(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.jobService.SubmitTest(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "test yuborishda xatolik", err.Error())
		return
	}

	response.Success(c, "test muvaffaqiyatli yuborildi", data)
}
