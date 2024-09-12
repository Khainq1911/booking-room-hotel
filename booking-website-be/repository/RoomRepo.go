package repository

import "booking-website-be/database"

type RoomRepo interface {
}

type RoomSql struct {
	Sql *database.Sql
}

func NewRoomRepo(sql *database.Sql) RoomRepo {
	return &RoomSql{
		Sql: sql,
	}
}
