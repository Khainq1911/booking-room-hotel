package repository

import (
	"booking-website-be/database"
	"booking-website-be/model"
	"context"
	"fmt"
	"time"
)

type AdminRepo interface {
	// rooms
	AddInforRoom(ctx context.Context, infor model.Room) (model.Room, error)
	DeleteRoomRepo(ctx context.Context, room_id int) error
	UpdateRoomInforRepo(ctx context.Context, room_id int, room model.RoomUpdate) error
	// bookings
	GetBookingsListRepo(ctx context.Context) ([]model.BookingList, error)
	GetDetailBookingRepo(ctx context.Context, booking_id int) ([]model.BookingList, error)
	UpdateBookingRepo(ctx context.Context, booking_id int, data model.BookingUpdate) error
	CancelBookingRepo(ctx context.Context, booking_id int) error
}

type AdminRepoDb struct {
	Sql *database.Sql
}

func NewAdminRepo(sql *database.Sql) AdminRepo {
	return &AdminRepoDb{
		Sql: sql,
	}
}

// booking action
func (db *AdminRepoDb) GetBookingsListRepo(ctx context.Context) ([]model.BookingList, error) {
	data := []model.BookingList{}
	query := `select * from bookings`

	if err := db.Sql.Db.Select(&data, query); err != nil {
		fmt.Println("failed to get booking list(GetBookingsListRepo)")
		return []model.BookingList{}, err
	}

	return data, nil
}

func (db *AdminRepoDb) GetDetailBookingRepo(ctx context.Context, booking_id int) ([]model.BookingList, error) {
	data := []model.BookingList{}

	query := `select * from bookings where booking_id = $1`
	if err := db.Sql.Db.Select(&data, query, booking_id); err != nil {
		fmt.Println("failed to get booking with id", booking_id, err)
		return []model.BookingList{}, err
	}

	return data, nil
}

func (db *AdminRepoDb) UpdateBookingRepo(ctx context.Context, booking_id int, data model.BookingUpdate) error {
	updateAt := time.Now()
	query := `update bookings
			set room_id = $1,
				check_in_date = $2,
				check_out_date = $3,
				total_price = $4,
				booking_status = $5,
				update_at = $6
				where booking_id = $7`

	result, err := db.Sql.Db.Exec(query, data.Room_id, data.CheckInDate, data.CheckOutDate, data.TotalPrice, data.BookingStatus, updateAt, booking_id)
	if err != nil {
		fmt.Println("failed to execute query", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("failed to get rowsAffected")
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("there is no booking_id")
	}

	return nil
}

func (db *AdminRepoDb) CancelBookingRepo(ctx context.Context, booking_id int) error {
	query := `update bookings set booking_status = 'canceled' where booking_id = $1`

	_, err := db.Sql.Db.Exec(query, booking_id)
	if err != nil {
		fmt.Println("failed to update in database (CancelBookingRepo)", err)
		return err
	}

	return nil
}

// room action

func (db *AdminRepoDb) AddInforRoom(ctx context.Context, infor model.Room) (model.Room, error) {
	query := `insert into Rooms (room_type, description, price, room_status, max_guest, image_url) values ( $1, $2, $3, $4, $5, $6)`

	if _, err := db.Sql.Db.Exec(query, infor.Room_type, infor.Description, infor.Price, infor.Room_status, infor.Max_guest, infor.Image_url); err != nil {
		fmt.Println("loi khi chen thong tin phong vao database(repository)")
		return model.Room{}, err
	}

	return infor, nil
}

func (db *AdminRepoDb) DeleteRoomRepo(ctx context.Context, room_id int) error {

	query := `delete from rooms where room_id = $1`

	if _, err := db.Sql.Db.Exec(query, room_id); err != nil {
		fmt.Println("failed to delete room from database", err)
		return err

	}

	return nil
}

func (db *AdminRepoDb) UpdateRoomInforRepo(ctx context.Context, room_id int, room model.RoomUpdate) error {

	query := `update rooms 
			set 
			room_type = $1,
			description = $2,
			price = $3,
			room_status = $4,
			max_guest = $5,
			image_url = $6
			where room_id = $7
			`
	result, err := db.Sql.Db.Exec(query, room.Room_type, room.Description, room.Price, room.Room_status, room.Max_guest, room.Image_url, room_id)
	if err != nil {
		fmt.Println("failed to update room information.", err)
		return err
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected Error", err)
		return err
	}

	if RowsAffected == 0 {
		return fmt.Errorf("room_id does not exist")
	}

	return nil
}

/*
User Authentication:
POST /register: Đăng ký người dùng mới. (done)
POST /login: Đăng nhập và nhận token xác thực. (done)
Room Management:
GET /rooms: Lấy danh sách các phòng. (done)
GET /rooms/:id: Lấy chi tiết thông tin phòng theo ID. (done)
Booking Management:
POST /bookings: Tạo mới một booking. (done)
GET /bookings: Lấy danh sách các booking của người dùng.(done)
GET /bookings/:id: Lấy chi tiết một booking.(done)
PUT /bookings/:id/cancel: Hủy một booking. (done)
Payment Management (tuỳ chọn):
POST /payments: Xử lý thanh toán.
GET /payments/:id: Lấy chi tiết thanh toán.
*/
/*
PUT /api/admin/rooms/:id (admin)
DELETE /api/admin/rooms/:id (admin) */
