package router

import (
	customerHandler "booking-website-be/handler"
	"booking-website-be/middleware"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo    *echo.Echo
	Handler customerHandler.CustomerHandler
}

func (api *Api) SetupRouter() {
	api.Echo.POST("/signUp", api.Handler.SaveAccount)
	api.Echo.POST("/signIn", api.Handler.CheckLogin)
	protected := api.Echo.Group("/")
	protected.Use(middleware.AuthenticateMiddleware)
	protected.GET("/", api.Handler.Home)
}
