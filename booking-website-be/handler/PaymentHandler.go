package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	Repo repository.PaymentRepo
}

func (u *PaymentHandler) CreatePayment(ctx echo.Context) error {
	req := model.CreatePayment{}
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.Repo.CreatePaymentRepo(ctx.Request().Context(), req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

func (u *PaymentHandler) ViewListPayment(ctx echo.Context) error {
	data, err := u.Repo.ViewListPaymentRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view data",
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u *PaymentHandler) ViewDetailPayment(ctx echo.Context) error {
	payment_id := ctx.Param("payment_id")

	data, err := u.Repo.ViewDetailPaymentRepo(ctx.Request().Context(), payment_id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u *PaymentHandler) UpdatePayment(ctx echo.Context) error {
	req := model.UpdatePayment{}
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	payment_id := ctx.Param("payment_id")

	err := u.Repo.UpdatePaymentRepo(ctx.Request().Context(), payment_id, req)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to view data",
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}
