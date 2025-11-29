package middleware

import (
	"geteway-service/internal/response"
	"geteway-service/internal/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header topilmadi", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, http.StatusUnauthorized, "Token formati noto‘g‘ri", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]
		secret := util.LoadEnv("JWT_ACCESS_SECRET")
		if secret == "" {
			response.Error(c, http.StatusInternalServerError, "JWT_SECRET topilmadi", nil)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response.Error(c, http.StatusUnauthorized, "Token yaroqsiz yoki muddati o‘tgan", err.Error())
			c.Abort()
			return
		}

		// Token valid bo‘lsa — foydalanuvchini context’ga qo‘shamiz
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if username, ok := claims["username"].(string); ok {
				c.Set("username", username)
			}
			if userID, ok := claims["user_id"].(string); ok {
				c.Set("user_id", userID)
			}
		}

		c.Next()
	}
}
