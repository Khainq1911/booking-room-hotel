package handler

import (
	"booking-website-be/model"
	"booking-website-be/repository"
	"fmt"
	"net/http"

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
