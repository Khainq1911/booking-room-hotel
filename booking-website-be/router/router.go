package router

import (
	"booking-website-be/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo            *echo.Echo
	CustomerHandler handler.CustomerHandler
	AdminHandler    handler.AdminHandler
	AccountHandler  handler.AccountHandler
}

func (api *Api) SetupRouter() {

	api.Echo.POST("customer/create", api.AccountHandler.CreateCustomer)
	api.Echo.GET("customers", api.AccountHandler.ViewCusList)
	api.Echo.GET("customers/:customer_id", api.AccountHandler.ViewCusDetail)
	api.Echo.PUT("customers/update/:customer_id", api.AccountHandler.UpdateCus)
	api.Echo.PUT("customers/delete/:customer_id", api.AccountHandler.DeleteCus)

	api.Echo.POST("employee/create", api.AccountHandler.CreateEmployee)
}
