package handler

import (
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type FilterHandler struct {
	filterService *service.FilterService
}

func NewFilterHandler(filterService *service.FilterService) *FilterHandler {
	return &FilterHandler{filterService: filterService}
}

func (h *FilterHandler) GetUsers(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.filterService.GetUsers(c, body)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "foydalanuvchilarni olishda xatolik", err.Error())
		return
	}

	response.Success(c, "foydalanuvchilar muvaffaqiyatli olindi", data)
}
