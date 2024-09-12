package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type TypeRoomRepo interface {
	AddTypeRoomRepo(ctx context.Context, typeRoom model.TypeRoom) error
	ViewtypeRoomRepo(ctx context.Context) ([]model.SelectTypeRoom, error)
	ViewDetailtypeRoomRepo(ctx context.Context, typeId string) ([]model.SelectTypeRoom, error)
	UpdateTypeRoomRepo(ctx context.Context, data model.UpdateTypeRoom, typeId string) error
	DeleteTypeRoomRepo(ctx context.Context, typeId string, typeRoom model.DeleteTypeRoom) error
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
	discount from typeroom `

	if err := db.Sql.Db.Select(&data, query); err != nil {
		return []model.SelectTypeRoom{}, err
	}

	return data, nil
}

// view detail type room
func (db *Sql) ViewDetailtypeRoomRepo(ctx context.Context, typeId string) ([]model.SelectTypeRoom, error) {
	data := []model.SelectTypeRoom{}

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
		return []model.SelectTypeRoom{}, err
	}

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
	where type_id = $3`

	current := time.Now()

	result, err := db.Sql.Db.Exec(query, current, typeRoom.DeleteBy, typeId)
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

/* func (db *CustomerRepoDb) FilterRoomRepo(ctx context.Context, room_type string, max_guest string, timeIn string, timeOut string) ([]model.Room, error) {
	var rooms []model.Room
	num := 1
	query := `SELECT * FROM rooms WHERE 1=1`
	params := []interface{}{}

	if room_type != "all" && room_type != "" {
		query += (` AND room_type = $` + strconv.Itoa(num))
		params = append(params, room_type)
		num += 1
	}

	if max_guest != "all" && max_guest != "" {
		query += (` AND max_guest <= $` + strconv.Itoa(num))
		params = append(params, max_guest)
		num += 1
	}

	if timeIn != "all" && timeOut != "all" && timeIn != "" && timeOut != "" {
		query += ` AND room_id NOT IN (
			SELECT room_id FROM bookings
			WHERE ($` + strconv.Itoa(num) + ` <= check_out_date AND $` + strconv.Itoa(num+1) + ` >= check_in_date)
		)`
		params = append(params, timeOut, timeIn)
		num += 2
	}

	if err := db.sql.Db.Select(&rooms, query, params...); err != nil {
		fmt.Println("Failed to filter data", err)
		return nil, err
	}

	return rooms, nil
} */
