package handler

import (
	"encoding/json"
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExamHandler struct {
	examService *service.ExamService
}

// NewExamHandler — yangi handler yaratish
func NewExamHandler(examService *service.ExamService) *ExamHandler {
	return &ExamHandler{examService: examService}
}

// CREATE — yangi exam yaratish
func (h *ExamHandler) Create(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.examService.CreateExam(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "exam yaratilmadi", err.Error())
		return
	}

	response.Success(c, "exam yaratildi", data)
}

// GET ALL — barcha examlar
func (h *ExamHandler) GetAll(c *gin.Context) {
	data, err := h.examService.GetExamList(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "examlarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "examlar ro‘yxati", data)
}

// GET BY ID — id orqali exam olish
func (h *ExamHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.examService.GetExamByID(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "examni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "exam topildi", data)
}

// UPDATE — examni yangilash
func (h *ExamHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID noto‘g‘ri formatda", err.Error())
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(body, &dataMap); err != nil {
		response.Error(c, http.StatusBadRequest, "json parse xato", err.Error())
		return
	}

	dataMap["id"] = id
	newBody, _ := json.Marshal(dataMap)

	data, err := h.examService.UpdateExam(c, newBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "exam yangilanmadi", err.Error())
		return
	}

	response.Success(c, "exam yangilandi", data)
}

// DELETE — examni o‘chirish
func (h *ExamHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.examService.DeleteExam(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "examni o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "exam o‘chirildi", data)
}
