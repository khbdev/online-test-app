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

type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler — yangi handler yaratish
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CREATE — yangi user yaratish
func (h *UserHandler) Create(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.userService.CreateUser(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "user yaratilmadi", err.Error())
		return
	}

	response.Success(c, "user yaratildi", data)
}

// GET ALL — barcha userlar
func (h *UserHandler) GetAll(c *gin.Context) {
	data, err := h.userService.GetAllUsers(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "userlarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "userlar ro‘yxati", data)
}

// GET BY ID — id orqali user olish
func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.userService.GetUserByID(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "userni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "user topildi", data)
}

// UPDATE — userni yangilash
func (h *UserHandler) Update(c *gin.Context) {
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

	data, err := h.userService.UpdateUser(c, newBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "user yangilanmadi", err.Error())
		return
	}

	response.Success(c, "user yangilandi", data)
}

// DELETE — userni o‘chirish
func (h *UserHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "id noto‘g‘ri formatda", err.Error())
		return
	}

	data, err := h.userService.DeleteUser(c, id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "userni o‘chirishda xatolik", err.Error())
		return
	}

	response.Success(c, "user o‘chirildi", data)
}
