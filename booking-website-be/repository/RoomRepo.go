package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type RoomRepo interface {
	AddRoomRepo(ctx context.Context, room model.AddRoom) error
	UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom) error
	DeleteRoomRepo(ctx context.Context, roomId string, Room model.DeleteRoom) error
	ViewListRoomRepo(ctx context.Context) ([]model.Room, error)
	ViewDetailRoomRepo(ctx context.Context, roomId string) ([]model.Room, error)
}

type RoomSql struct {
	Sql *database.Sql
}

func NewRoomRepo(sql *database.Sql) RoomRepo {
	return &RoomSql{
		Sql: sql,
	}
}

// add type room
func (db *RoomSql) AddRoomRepo(ctx context.Context, room model.AddRoom) error {
	query := `insert into room (
	room_name,
	type_id, 
	floor,
	status,
	price_override,
	cleaning_status,
	check_in_time,
	check_out_time,
	current_guest,
	note, 
	createtime,
	createby) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, room.RoomName, room.TypeID, room.Floor, room.Status, room.PriceOverride,
		room.CleaningStatus, room.CheckInTime, room.CheckOutTime, room.CurrentGuest, room.Note, current, room.CreateBy); err != nil {
		return err
	}

	return nil
}

// view list room
func (db *RoomSql) ViewListRoomRepo(ctx context.Context) ([]model.Room, error) {
	data := []model.Room{}

	query := `select room_id,
	room_name,
	type_id, 
	floor,
	status,
	price_override,
	cleaning_status,
	check_in_time,
	check_out_time,
	current_guest,
	note from room`

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.Room{}, err
	}

	return data, nil
}

// view detail room
func (db *RoomSql) ViewDetailRoomRepo(ctx context.Context, roomId string) ([]model.Room, error) {
	data := []model.Room{}
	query := `select 
	room_name,
	type_id, 
	floor,
	status,
	price_override,
	cleaning_status,
	check_in_time,
	check_out_time,
	current_guest,
	note from room where room_id = $1`

	if err := db.Sql.Db.Select(&data, query, roomId); err != nil {
		return []model.Room{}, err
	}

	return data, nil
}

// update room
func (db *RoomSql) UpdateRoomRepo(ctx context.Context, roomId string, room model.UpdateRoom) error {
	query := `update room set
	room_name = $1,
	type_id = $2, 
	floor = $3,
	status = $4,
	price_override = $5,
	cleaning_status = $6,
	check_in_time = $7,
	check_out_time = $8,
	current_guest = $9,
	note = $10, 
	updatetime = $11,
	updateby = $12 where room_id = $13
	`
	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, room.RoomName, room.TypeID, room.Floor,
		room.Status, room.PriceOverride, room.CleaningStatus, room.CheckInTime, room.CheckOutTime,
		room.CurrentGuest, room.Note, current, room.UpdateBy, roomId); err != nil {
		return err
	}

	return nil
}

// delete room
func (db *RoomSql) DeleteRoomRepo(ctx context.Context, roomId string, Room model.DeleteRoom) error {
	query := `update room
	set deletetime = $1,
		deleteby = $2
	where room_id = $3`

	current := time.Now()

	result, err := db.Sql.Db.Exec(query, current, Room.DeleteBy, roomId)
	if err != nil {
		return err
	}

	rowwAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowwAffected == 0 {
		return fmt.Errorf("column is not exist")
	}

	return nil

}
