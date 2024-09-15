package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"strconv"
	"time"
)

type TypeRoomRepo interface {
	AddTypeRoomRepo(ctx context.Context, typeRoom model.TypeRoom) error
	ViewtypeRoomRepo(ctx context.Context) ([]model.SelectTypeRoom, error)
	ViewDetailtypeRoomRepo(ctx context.Context, typeId string) ([]model.SelectDetail, error)
	UpdateTypeRoomRepo(ctx context.Context, data model.UpdateTypeRoom, typeId string) error
	DeleteTypeRoomRepo(ctx context.Context, typeId string, typeRoom model.DeleteTypeRoom) error
	FilterTypeRoomRepo(ctx context.Context, TypeName string, max_occupancy string, timeIn string, timeOut string) ([]model.SelectTypeRoom, error)
}

type Sql struct {
	Sql *database.Sql
}

func NewTypeRoomRepo(sql *database.Sql) TypeRoomRepo {
	return &Sql{
		Sql: sql,
	}
}

// add type room
func (db *Sql) AddTypeRoomRepo(ctx context.Context, typeRoom model.TypeRoom) error {
	query := `insert into typeroom (
	type_name,
	description,
	price_per_night,
	max_occupancy,
	room_size,
	image_url,
	status,
	discount,
	createtime,
	createby) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query,
		typeRoom.TypeName,
		typeRoom.Description,
		typeRoom.PricePerNight,
		typeRoom.MaxOccupancy,
		typeRoom.RoomSize,
		typeRoom.ImageURL,
		typeRoom.Status,
		typeRoom.Discount,
		current,
		typeRoom.CreateBy); err != nil {
		return err
	}

	return nil
}

// view type room
func (db *Sql) ViewtypeRoomRepo(ctx context.Context) ([]model.SelectTypeRoom, error) {
	data := []model.SelectTypeRoom{}

	query := `select type_id,
	type_name,
	description,
	price_per_night,
	max_occupancy,
	room_size,
	image_url,
	status,
	discount from typeroom where is_deleted = $1`

	if err := db.Sql.Db.Select(&data, query, false); err != nil {
		return []model.SelectTypeRoom{}, err
	}

	return data, nil
}

// view detail type room
func (db *Sql) ViewDetailtypeRoomRepo(ctx context.Context, typeId string) ([]model.SelectDetail, error) {
	data := []model.SelectDetail{}

	data2 := []model.Room{}

	query := `select type_id,
	type_name,
	description,
	price_per_night,
	max_occupancy,
	room_size,
	image_url,
	status,
	discount from typeroom where type_id = $1`

	if err := db.Sql.Db.Select(&data, query, typeId); err != nil {
		return []model.SelectDetail{}, err
	}

	query2 := `select room_id, room_name, type_id, floor, status, price_override, cleaning_status, check_in_time, check_out_time, current_guest, note from room where type_id = $1`
	if err := db.Sql.Db.Select(&data2, query2, typeId); err != nil {
		return []model.SelectDetail{}, err
	}

	data[0].Rooms = data2

	return data, nil
}

// update type room
func (db *Sql) UpdateTypeRoomRepo(ctx context.Context, data model.UpdateTypeRoom, typeId string) error {
	query := ` update typeroom
	set 
	type_name = $1,
	description = $2,
	price_per_night = $3,
	max_occupancy = $4,
	room_size = $5,
	image_url = $6,
	status = $7,
	discount = $8,
	updatetime = $9,
	updateby = $10
	where type_id = $11`

	current := time.Now()

	if _, err := db.Sql.Db.Exec(query, data.TypeName, data.Description, data.PricePerNight, data.MaxOccupancy,
		data.RoomSize, data.ImageURL, data.Status, data.Discount, current, data.UpdateBy, typeId); err != nil {
		return err
	}

	return nil
}

// delete room type
func (db *Sql) DeleteTypeRoomRepo(ctx context.Context, typeId string, typeRoom model.DeleteTypeRoom) error {
	query := `update typeroom
	set deletetime = $1,
		deleteby = $2
		is_deleted = $3
	where type_id = $4`

	current := time.Now()
	isDeleted := false
	result, err := db.Sql.Db.Exec(query, current, typeRoom.DeleteBy, isDeleted, typeId)
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

func (db *Sql) FilterTypeRoomRepo(ctx context.Context, TypeName string, max_occupancy string, timeIn string, timeOut string) ([]model.SelectTypeRoom, error) {
	typeRoom := []model.SelectTypeRoom{}
	num := 1
	query := `SELECT 
    type_id,
    type_name,
    description,
    price_per_night,
    max_occupancy,
    room_size,
    image_url,
    status,
    discount 
FROM 
    typeroom `
	params := []interface{}{}

	if TypeName != "all" && TypeName != "" {
		query += (` AND typename = $` + strconv.Itoa(num))
		params = append(params, TypeName)
		num += 1
	}

	if max_occupancy != "all" && max_occupancy != "" {
		query += (` AND max_occupancy <= $` + strconv.Itoa(num))
		params = append(params, max_occupancy)
		num += 1
	}

	if timeIn != "all" && timeOut != "all" && timeIn != "" && timeOut != "" {
		query += ` AND type_id NOT IN (
			SELECT type_id FROM room
			WHERE ($` + strconv.Itoa(num) + ` <= check_out_time AND $` + strconv.Itoa(num+1) + ` >= check_in_time)
		)`
		params = append(params, timeOut, timeIn)
		num += 2
	}

	fmt.Println(query)

	if err := db.Sql.Db.Select(&typeRoom, query, params...); err != nil {
		return nil, err
	}

	return typeRoom, nil
}
