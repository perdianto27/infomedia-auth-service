package auth

import (
	"main/helpers"
	"main/models/base"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RefreshTokenController(c echo.Context) error {

	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind(&request); err != nil || request.RefreshToken == "" {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "refresh_token is required",
			Data:    nil,
		})
	}

	token, err := jwt.ParseWithClaims(request.RefreshToken, &helpers.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_REFRESH_JWT")), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, base.BaseResponse{
			Status:  false,
			Message: "Invalid or expired refresh token",
			Data:    nil,
		})
	}

	claims := token.Claims.(*helpers.RefreshClaims)

	// Generate Access Token baru
	newAccessToken, err := helpers.GenerateTokenJWT(claims.Email, "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed generate new access token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Token refreshed successfully",
		Data: map[string]string{
			"access_token": newAccessToken,
		},
	})
}