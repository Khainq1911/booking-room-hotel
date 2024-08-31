package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type CustomerRepo interface {
	GetAllRoomRepo(ctx context.Context) ([]model.Room, error)
	SelectRoomRepo(ctx context.Context, room_id int) ([]model.Room, error)
	BookingRoomRepo(ctx context.Context, booking model.BookingRoom) (model.BookingRoom, error)
}

type CustomerRepoDb struct {
	sql *database.Sql
}

func NewCustomerRepo(sql *database.Sql) CustomerRepo {
	return &CustomerRepoDb{
		sql: sql,
	}
}

func (db *CustomerRepoDb) GetAllRoomRepo(ctx context.Context) ([]model.Room, error) {
	data := []model.Room{}
	query := `SELECT * FROM Rooms`

	if err := db.sql.Db.Select(&data, query); err != nil {
		fmt.Println(" loi lay phong tu database", err)
		return []model.Room{}, err
	}

	return data, nil
}

func (db *CustomerRepoDb) SelectRoomRepo(ctx context.Context, room_id int) ([]model.Room, error) {
	data := []model.Room{}
	query := `select * from rooms where room_id = $1`

	if err := db.sql.Db.Select(&data, query, room_id); err != nil {
		fmt.Println("loi lay phong tu database (SelectRoomRepo); ", err)
		return []model.Room{}, err
	}

	return data, nil
}

func (db *CustomerRepoDb) BookingRoomRepo(ctx context.Context, booking model.BookingRoom) (model.BookingRoom, error) {
	query := `insert into bookings(room_id, user_id, check_in_date, check_out_date, total_price, booking_status, create_at, update_at) 
			values ($1, $2, $3, $4, $5, $6, $7, $8)`

	create_at := time.Now()
	update_at := time.Now()

	_, err := db.sql.Db.Exec(query,
		booking.Room_id,
		booking.User_id,
		booking.CheckInDate,
		booking.CheckOutDate,
		booking.TotalPrice,
		booking.BookingStatus,
		create_at, update_at)
	if err != nil {
		fmt.Println("failed to insert data to database at BookingRoomRepo")
		return model.BookingRoom{}, err
	}

	return booking, nil
}
