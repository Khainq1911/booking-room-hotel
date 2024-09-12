package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TypeRoomHandler struct {
	TypeRoomRepo repository.TypeRoomRepo
}

func (u *TypeRoomHandler) AddTypeRoom(ctx echo.Context) error {
	req := model.TypeRoom{}

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.TypeRoomRepo.AddTypeRoomRepo(ctx.Request().Context(), req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to add data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}

// view type room
func (u *TypeRoomHandler) ViewTypeRoom(ctx echo.Context) error {

	data, err := u.TypeRoomRepo.ViewtypeRoomRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to add data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

// view detail type room
func (u *TypeRoomHandler) ViewDetailTypeRoom(ctx echo.Context) error {
	typeId := ctx.Param("type_id")

	data, err := u.TypeRoomRepo.ViewDetailtypeRoomRepo(ctx.Request().Context(), typeId)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to add data",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

//update type room

func (u *TypeRoomHandler) UpdateTypeRoom(ctx echo.Context) error {
	typeId := ctx.Param("type_id")

	req := model.UpdateTypeRoom{}
	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.TypeRoomRepo.UpdateTypeRoomRepo(ctx.Request().Context(), req, typeId); err != nil {
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

func (u *TypeRoomHandler) DeleteTypeRoom(ctx echo.Context) error {
	var req model.DeleteTypeRoom

	if err := ctx.Bind(&req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	type_id := ctx.Param("type_id")

	if err := u.TypeRoomRepo.DeleteTypeRoomRepo(ctx.Request().Context(), type_id, req); err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update type room",
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       req,
	})
}
