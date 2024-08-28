package model

type BookingRoom struct {
	/* Booking_id    int       `json:"booking_id" db:"booking_id"` */
	Room_id int `json:"room_id" db:"room_id"`
	User_id int `json:"user_id" db:"user_id"`
	/* CheckInDate   time.Time `json:"check_in_date" db:"check_in_date"`
	CheckOutDate  time.Time `json:"check_out_date" db:"check_out_date"` */
	TotalPrice    float32 `json:"total_price" db:"total_price"`
	BookingStatus string  `json:"booking_status" db:"booking_status"`
}
