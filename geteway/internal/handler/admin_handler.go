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

type AdminHandler struct {
	adminService *service.AdminService
}

// NewAdminHandler — yangi handler yaratish
func NewAdminHandler(adminService *service.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

// CREATE — yangi admin yaratish
func (h *AdminHandler) Create(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.adminService.CreateAdmin(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "admin yaratilmadi", err.Error())
		return
	}

	response.Success(c, "admin yaratildi", data)
}

// GET ALL — barcha adminlar
func (h *AdminHandler) GetAll(c *gin.Context) {
	data, err := h.adminService.GetAdminList(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "adminlarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "adminlar ro‘yxati", data)
}

// GET BY ID — id orqali adminni olish
func (h *AdminHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.adminService.GetAdminByID(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "adminni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "admin topildi", data)
}

// UPDATE — adminni yangilash
func (h *AdminHandler) Update(c *gin.Context) {
	// 🧩 URL dan id olish
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "ID noto‘g‘ri formatda", err.Error())
		return
	}

	// 🧩 Body ni o‘qish
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	// 🧩 JSON body ni map ga parse qilish
	var dataMap map[string]interface{}
	if err := json.Unmarshal(body, &dataMap); err != nil {
		response.Error(c, http.StatusBadRequest, "json parse xato", err.Error())
		return
	}

	// 🧩 ID ni body ichiga qo‘shamiz
	dataMap["id"] = id

	// 🧩 Yangi body tayyorlaymiz (id qo‘shilgan)
	newBody, _ := json.Marshal(dataMap)

	// 🧩 Servisga yuboramiz
	data, err := h.adminService.UpdateAdmin(c, newBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "admin yangilanmadi", err.Error())
		return
	}

	// ✅ Muvaffaqiyatli javob
	response.Success(c, "admin yangilandi", data)
}


// DELETE — adminni o‘chirish
func (h *AdminHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.adminService.DeleteAdmin(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "adminni o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "admin o‘chirildi", data)
}
