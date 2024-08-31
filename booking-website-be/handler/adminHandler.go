package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Repo repository.AdminRepo
}

func (u AdminHandler) AddInforRoomHandler(ctx echo.Context) error {
	data := model.Room{}

	if err := ctx.Bind(&data); err != nil {
		fmt.Println("loi khong the nhan du lieu da chen (AddInforRoomHandler)")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "loi khi nhan du lieu tu request",
		})
	}

	roomInfor, err := u.Repo.AddInforRoom(ctx.Request().Context(), data)
	if err != nil {
		fmt.Println("loi khi chen du lieu o AddInforRoomHandler")
		return ctx.JSON(http.StatusConflict, echo.Map{
			"message": "loi khi chen du lieu o handler",
		})
	}

	return ctx.JSON(http.StatusAccepted, echo.Map{
		"message": "succesful",
		"data":    roomInfor,
	})
}

func (u AdminHandler) GetBookingList(ctx echo.Context) error {

	data, err := u.Repo.GetBookingsListRepo(ctx.Request().Context())
	if err != nil {
		fmt.Println("error in database")
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get bookings list",
			Data:       nil,
		})
	}
	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u AdminHandler) GetDetailBooking(ctx echo.Context) error {
	booking_id, err := strconv.Atoi(ctx.QueryParam("booking_id"))

	if err != nil {
		fmt.Println("failed to get booking_id")
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get params",
			Data:       nil,
		})
	}

	data, err := u.Repo.GetDetailBookingRepo(ctx.Request().Context(), booking_id)
	if err != nil {
		fmt.Println("failed in database")
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed in database",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u *AdminHandler) CancelBooking(ctx echo.Context) error {

	booking_id, err := strconv.Atoi(ctx.QueryParam("booking_id"))
	if err != nil {
		fmt.Println("failed to get booking_id")
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get params in cancel booking",
			Data:       nil,
		})
	}

	if err := u.Repo.CancelBookingRepo(ctx.Request().Context(), booking_id); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed in database",
			Data:       nil,
		})
	}

	data, err := u.Repo.GetDetailBookingRepo(ctx.Request().Context(), booking_id)
	if err != nil {
		fmt.Println("failed in database")
		return ctx.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed in database",
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "successful",
		Data:       data,
	})
}

func (u *AdminHandler) DeleteRoom(ctx echo.Context) error {
	room_id, err := strconv.Atoi(ctx.Param("room_id"))
	if err != nil {
		fmt.Println("failed to get room_id")
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get room_id",
		})
	}

	if err := u.Repo.DeleteRoomRepo(ctx.Request().Context(), room_id); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to delete rooms",
		})
	}

	return ctx.JSON(http.StatusOK, model.ResWithOutData{
		StatusCode: http.StatusOK,
		Message:    "successful",
	})
}

func (u *AdminHandler) UpdateRoom(ctx echo.Context) error {
	data := model.RoomUpdate{}

	room_id, err := strconv.Atoi(ctx.Param("room_id"))
	if err != nil {
		fmt.Println("failed to get room_id")
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get room_id",
		})
	}

	if err := ctx.Bind(&data); err != nil {
		fmt.Println("failed to bind data")
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to bind data",
		})
	}

	if err := u.Repo.UpdateRoomInforRepo(ctx.Request().Context(), room_id, data); err != nil {

		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, model.ResWithOutData{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to update",
		})

	}
	fmt.Println(data, room_id)
	return ctx.JSON(http.StatusOK, model.ResWithOutData{
		StatusCode: http.StatusOK,
		Message:    "successful",
	})
}
