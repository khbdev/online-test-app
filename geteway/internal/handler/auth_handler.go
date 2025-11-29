package handler

import (
	"geteway-service/internal/response"
	"geteway-service/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// ✅ Login
func (h *AuthHandler) Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.authService.Login(c, body)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "login amalga oshmadi", err.Error())
		return
	}

	response.Success(c, "login muvaffaqiyatli", data)
}

// ✅ Refresh token orqali yangi access token olish
func (h *AuthHandler) Refresh(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "bodyni o‘qishda xatolik", err.Error())
		return
	}

	data, err := h.authService.Refresh(c, body)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "refresh amalga oshmadi", err.Error())
		return
	}

	response.Success(c, "token yangilandi", data)
}
