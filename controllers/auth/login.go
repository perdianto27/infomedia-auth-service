package auth

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"main/configs"
	"main/models/base"
	usermodel "main/models/user"
	helpers "main/helpers"
)

func LoginController(c echo.Context) error {
	var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
	}

	if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, base.BaseResponse{
					Status:  false,
					Message: "Invalid request format",
					Data:    nil,
			})
	}
	fmt.Println("request", request);
	if request.Email == "" || request.Password == "" {
			return c.JSON(http.StatusBadRequest, base.BaseResponse{
					Status:  false,
					Message: "Login dan Password wajib diisi",
					Data:    nil,
			})
	}

	var user usermodel.User
	if err := configs.DB.
			Where("username = ? OR email = ?", request.Email, request.Email).
			First(&user).Error; err != nil {

			return c.JSON(http.StatusUnauthorized, base.BaseResponse{
					Status:  false,
					Message: "User tidak ditemukan",
					Data:    nil,
			})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, base.BaseResponse{
					Status:  false,
					Message: "Password salah",
					Data:    nil,
			})
	}

	token, err := helpers.GenerateTokenJWT(user.Email, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed generate token",
			Data:    nil,
		})
	}


	refreshToken, err := helpers.GenerateRefreshToken(user.Email)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, base.BaseResponse{
					Status:  false,
					Message: "Failed generate refresh token",
					Data:    nil,
			})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
			Status:  true,
			Message: "Login berhasil",
			Data: echo.Map{
					"access_token": token,
					"refresh_token": refreshToken,
					"user": echo.Map{
							"id":       user.Id,
							"username": user.Username,
							"email":    user.Email,
							"name":     user.Name,
					},
			},
	})
}