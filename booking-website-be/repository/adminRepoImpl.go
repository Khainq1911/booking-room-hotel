package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
)

type AdminRepo interface {
	AddInforRoom(ctx context.Context, infor model.Room) (model.Room, error)
}

type AdminRepoDb struct {
	Sql *database.Sql
}

func NewAdminRepo(sql *database.Sql) AdminRepo {
	return &AdminRepoDb{
		Sql: sql,
	}
}

func (db *AdminRepoDb) AddInforRoom(ctx context.Context, infor model.Room) (model.Room, error) {
	query := `insert into Rooms (room_type, description, price, room_status, max_guest, image_url) values ( $1, $2, $3, $4, $5, $6)`

	if _, err := db.Sql.Db.Exec(query, infor.Room_type, infor.Description, infor.Price, infor.Room_status, infor.Max_guest, infor.Image_url); err != nil {
		fmt.Println("loi khi chen thong tin phong vao database(repository)")
		return model.Room{}, err
	}

	return infor, nil
}
