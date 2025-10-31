package user

import (
	"main/configs"
	"main/models/base"
	usermodel "main/models/user"
	helpers "main/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AddUserController(c echo.Context) error {
	var user usermodel.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid request data",
			Data:    nil,
		})
	}

	if err := configs.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error; err == nil {
		return c.JSON(http.StatusConflict, base.BaseResponse{
			Status:  false,
			Message: "User already exists",
			Data:    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed encrypt password",
			Data:    nil,
		})
	}

	user.Password = string(hashedPassword)

	result := configs.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}
	
	go helpers.SendRegisterSuccessEmail(user.Email, user.Name)

	return c.JSON(http.StatusCreated, base.BaseResponse{
		Status:  true,
		Message: "Success register",
		Data:    user,
	})
}
