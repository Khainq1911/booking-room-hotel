package main

import (
	"booking-website-be/database"
	"booking-website-be/handler"
	"booking-website-be/repository"

	"booking-website-be/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	sql := &database.Sql{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Port:     5432,
		Dbname:   "booking-room-hotel",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.CORS())

	TypeRoomDb := handler.TypeRoomHandler{
		TypeRoomRepo: repository.NewTypeRoomRepo(sql),
	}

	AccountHandler := handler.AccountHandler{
		Repo: repository.NewAccountRepo(sql),
	}

	RoomDb := handler.RoomHandler{
		RoomRepo: repository.NewRoomRepo(sql),
	}
	api := router.Api{
		Echo:            e,
		AccountHandler:  AccountHandler,
		TypeRoomHandler: TypeRoomDb,
		RoomHandler:     RoomDb,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1912"))
}
