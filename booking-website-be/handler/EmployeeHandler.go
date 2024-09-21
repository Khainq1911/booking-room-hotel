package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"booking-website-be/security"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	EmployeeRepo repository.EmployeeRepo
}

//EMPLOYEE

// create
func (u *EmployeeHandler) CreateEmployee(ctx echo.Context) error {
	req := model.CreateEmp{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.EmployeeRepo.CreateEmpRepo(ctx.Request().Context(), req); err != nil {
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

//view list

func (u *EmployeeHandler) ViewListEmp(ctx echo.Context) error {

	data, err := u.EmployeeRepo.ViewListEmpRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "succesful",
		Data:       data,
	})
}

// view detail
func (u *EmployeeHandler) ViewDetailEmp(ctx echo.Context) error {

	employee_id := ctx.Param("employee_id")

	data, err := u.EmployeeRepo.ViewDetailEmpRepo(ctx.Request().Context(), employee_id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "succesful",
		Data:       data,
	})
}

//update employee

func (u *EmployeeHandler) UpdateEmp(ctx echo.Context) error {
	req := model.UpdateEmp{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	employee_id := ctx.Param("employee_id")

	if err := u.EmployeeRepo.UpdateEmpRepo(ctx.Request().Context(), employee_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

// delete
func (u *EmployeeHandler) DeleteEmp(ctx echo.Context) error {
	req := model.DeleteEmp{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	employee_id := ctx.Param("employee_id")

	if err := u.EmployeeRepo.DeleteEmpRepo(ctx.Request().Context(), employee_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to delete data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

func (u *EmployeeHandler) CheckLogin(ctx echo.Context) error {
	req := model.User{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	data, err := u.EmployeeRepo.CheckLogin(ctx.Request().Context(), req.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to delete data",
		})
	}

	if req.PassWord != data[0].PassWord {
		return ctx.JSON(http.StatusUnauthorized, model.ResWithOutData{
			StatusCode: http.StatusUnauthorized,
			Message:    "failed to delete data",
		})
	}

	claims := jwt.MapClaims{
		"employee_id": data[0].EmployeeID,
		"is_admin":    data[0].IsAdmin,
		"exp":         time.Now().Add(time.Hour * 2).Unix(),
	}

	token, err := security.GenToken(&claims, ctx.Echo().AcquireContext())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to gentoken",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
