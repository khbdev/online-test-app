package handler

import (
	"encoding/json"
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TestSectionHandler struct {
	testService *service.TestService
}

func NewTestSectionHandler(testService *service.TestService) *TestSectionHandler {
	return &TestSectionHandler{testService: testService}
}

//
// ======================
// 🧩 SECTION CRUD
// ======================
//

func (h *TestSectionHandler) CreateSection(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.testService.CreateSection(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Section yaratishda xatolik", err.Error())
		return
	}

	response.Success(c, "Section muvaffaqiyatli yaratildi", data)
}

func (h *TestSectionHandler) GetSectionById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.GetSectionById(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Sectionni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "Section ma’lumotlari", data)
}

func (h *TestSectionHandler) GetAllSections(c *gin.Context) {
	data, err := h.testService.GetAllSections(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Sectionlarni olishda xatolik", err.Error())
		return
	}
	response.Success(c, "Sectionlar ro‘yxati", data)
}

func (h *TestSectionHandler) UpdateSection(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	// ✅ URL dan olingan id’ni body ichiga joylaymiz
	var req map[string]interface{}
	_ = json.Unmarshal(body, &req)
	req["id"] = id
	body, _ = json.Marshal(req)

	data, err := h.testService.UpdateSection(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Section yangilashda xatolik", err.Error())
		return
	}

	response.Success(c, "Section muvaffaqiyatli yangilandi", data)
}

func (h *TestSectionHandler) DeleteSection(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.DeleteSection(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Section o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "Section muvaffaqiyatli o‘chirildi", data)
}

//
// ======================
// ❓ QUESTION CRUD
// ======================
//

func (h *TestSectionHandler) CreateQuestion(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.testService.CreateQuestion(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Savol yaratishda xatolik", err.Error())
		return
	}

	response.Success(c, "Savol muvaffaqiyatli yaratildi", data)
}

func (h *TestSectionHandler) GetQuestionById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.GetQuestionById(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Savolni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "Savol ma’lumotlari", data)
}

func (h *TestSectionHandler) GetAllQuestions(c *gin.Context) {
	data, err := h.testService.GetAllQuestions(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Savollarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "Savollar ro‘yxati", data)
}

func (h *TestSectionHandler) UpdateQuestion(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	// ✅ URL dan olingan id’ni body ichiga joylaymiz
	var req map[string]interface{}
	_ = json.Unmarshal(body, &req)
	req["id"] = id
	body, _ = json.Marshal(req)

	data, err := h.testService.UpdateQuestion(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Savolni yangilashda xatolik", err.Error())
		return
	}

	response.Success(c, "Savol muvaffaqiyatli yangilandi", data)
}

func (h *TestSectionHandler) DeleteQuestion(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.DeleteQuestion(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Savolni o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "Savol muvaffaqiyatli o‘chirildi", data)
}

//
// ======================
// 🧩 OPTION CRUD
// ======================
//

func (h *TestSectionHandler) CreateOption(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.testService.CreateOption(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Option yaratishda xatolik", err.Error())
		return
	}

	response.Success(c, "Option muvaffaqiyatli yaratildi", data)
}

func (h *TestSectionHandler) GetOptionById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.GetOptionById(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Optionni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "Option ma’lumotlari", data)
}

func (h *TestSectionHandler) GetAllOptions(c *gin.Context) {
	data, err := h.testService.GetAllOptions(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Optionlarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "Optionlar ro‘yxati", data)
}

func (h *TestSectionHandler) UpdateOption(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	// ✅ URL dan olingan id’ni body ichiga joylaymiz
	var req map[string]interface{}
	_ = json.Unmarshal(body, &req)
	req["id"] = id
	body, _ = json.Marshal(req)

	data, err := h.testService.UpdateOption(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Optionni yangilashda xatolik", err.Error())
		return
	}

	response.Success(c, "Option muvaffaqiyatli yangilandi", data)
}

func (h *TestSectionHandler) DeleteOption(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	data, err := h.testService.DeleteOption(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Optionni o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "Option muvaffaqiyatli o‘chirildi", data)
}
