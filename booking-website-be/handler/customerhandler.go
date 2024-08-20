package handler

import (
	"booking-website-be/model"
	repo "booking-website-be/repository"
	"booking-website-be/secure"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	Repo repo.CustomerRepo
}

func (u *CustomerHandler) SaveAccount(ctx echo.Context) error {
	req := model.SaveAccount{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in bind data",
		})
	}

	hashedPassword, err := secure.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error in hashing password",
		})
	}

	user := model.SaveAccount{
		Name:     req.Name,
		Dob:      req.Dob,
		Phone:    req.Phone,
		Password: hashedPassword,
		Email:    req.Email,
		Role:     req.Role,
	}
	_, err = u.Repo.SaveAccountRepo(user, ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusConflict, echo.Map{
			"message": "error in database",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "successful",
		"data":    user,
	})
}

func (u *CustomerHandler) CheckLogin(ctx echo.Context) error {
	req := model.SignIn{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in bind data",
		})
	}

	fmt.Println(req)

	data, err := u.Repo.CheckSignInRepo(ctx.Request().Context(), req.Phone)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err,
		})
	}
	if (data == model.SignIn{}) {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "user not found",
		})
	}

	if !secure.CheckPassword(data.Password, req.Password) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid credentials",
		})
	}

	token, err := generateJWTToken(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error generating token",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "login successful",
		"token":   token,
	})

}
