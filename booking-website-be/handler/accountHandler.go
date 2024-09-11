package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"

	"booking-website-be/security"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	Repo repository.AccountRepo
}

func (u *AccountHandler) CheckLogin(ctx echo.Context) error {
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
	cookie.Expires = time.Now().Add(60 * time.Minute)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "login successful",
		"data":    data,
		"token":   token,
	})
}

func (u *AccountHandler) Logout(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Unix(0, 0)

	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, model.ResWithOutData{
		StatusCode: http.StatusOK,
		Message:    "succesful",
	})
}

// create customer handler
func (u *AccountHandler) CreateCustomer(ctx echo.Context) error {
	req := model.CreateCus{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.Repo.CreateCusRepo(ctx.Request().Context(), req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create customer",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})

}

// view list customers

func (u *AccountHandler) ViewCusList(ctx echo.Context) error {
	data, err := u.Repo.ViewCusListRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get data",
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

// view detail customer
func (u *AccountHandler) ViewCusDetail(ctx echo.Context) error {

	customer_id := ctx.Param("customer_id")

	data, err := u.Repo.ViewCusDetailRepo(ctx.Request().Context(), customer_id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view customer",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

// update customer
func (u *AccountHandler) UpdateCus(ctx echo.Context) error {
	req := model.UpdateCus{}
	customer_id := ctx.Param("customer_id")

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.Repo.UpdateCusRepo(ctx.Request().Context(), customer_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update customer",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

// delete customer
func (u *AccountHandler) DeleteCus(ctx echo.Context) error {
	var req model.DeleteCus

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	customer_id := ctx.Param("customer_id")

	if err := u.Repo.DeleteCusRepo(ctx.Request().Context(), customer_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update customer",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

//EMPLOYEE

// create
func (u *AccountHandler) CreateEmployee(ctx echo.Context) error {
	req := model.CreateEmp{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.Repo.CreateEmpRepo(ctx.Request().Context(), req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create employee",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}
