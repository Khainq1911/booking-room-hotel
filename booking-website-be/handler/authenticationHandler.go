package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"booking-website-be/security"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthenHandler struct {
	Repo repository.AuthenticationRepo
}

func (u *AuthenHandler) SaveAccount(ctx echo.Context) error {
	req := model.SaveAccount{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in bind data",
		})
	}

	hashedPassword, err := security.HashPassword(req.Password)
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

func (u *AuthenHandler) CheckLogin(ctx echo.Context) error {
	req := model.Request{}
	

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error in bind data",
		})
	}



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

	if !security.CheckPassword(data.Password, req.Password) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid credentials",
		})
	}

	token, err := security.GenerateJWTToken(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error generating token",
		})
	}

	
	cookie := new(http.Cookie) 
	cookie.Name = "token"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true                          
	cookie.Secure = true                           
	cookie.Expires = time.Now().Add(24 * time.Hour) 
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "login successful",
		"data":    data,
		"token":   token,
	})
}
