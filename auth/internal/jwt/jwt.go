package jwt

import (
	"auth-service/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

// --- Token juftligini yaratish ---
func GenerateTokens(username string) (*TokenPair, error) {
	now := time.Now()

	accessClaims := jwt.MapClaims{
		"sub":  username,
		"exp":  now.Add(config.JWT.AccessTTL).Unix(),
		"iat":  now.Unix(),
		"type": "access",
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(config.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"sub":  username,
		"exp":  now.Add(config.JWT.RefreshTTL).Unix(),
		"iat":  now.Unix(),
		"type": "refresh",
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(config.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}

// --- Refresh token validligini tekshir va yangi token juftligi yarat ---
func RefreshTokens(refreshToken string) (*TokenPair, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.JWT.RefreshSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	if claims["type"] != "refresh" {
		return nil, jwt.ErrTokenInvalidClaims
	}

	username, ok := claims["sub"].(string)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Agar refresh token expired bo‘lsa → login qilinish kerak
	exp := int64(claims["exp"].(float64))
	if time.Unix(exp, 0).Before(time.Now()) {
		return nil, errors.New("refresh token expired, please login again")
	}

	// ✅ Yangi token juftligini yarat
	return GenerateTokens(username)
}
