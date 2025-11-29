package handler

import (
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type TestHandler struct {
	testService *service.GenerateService
}

func NewGenerateHandler(testService *service.GenerateService) *TestHandler {
	return &TestHandler{testService: testService}
}

// GenerateTest — yangi test yaratish
func (h *TestHandler) GenerateTest(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "body o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.testService.GenerateTest(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "test yaratishda xatolik", err.Error())
		return
	}

	response.Success(c, "test muvaffaqiyatli yaratildi", data)
}

// GetTest — mavjud testni olish
func (h *TestHandler) GetTest(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "body o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.testService.GetTest(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "testni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "test muvaffaqiyatli olindi", data)
}
