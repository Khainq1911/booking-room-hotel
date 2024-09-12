package model

import "time"

type Room struct {
	RoomID         string     `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	TypeID         string     `json:"type_id" db:"type_id"`
	Floor          int        `json:"floor" db:"floor"`
	Status         string     `json:"status" db:"status"`
	PriceOverride  float64    `json:"price_override" db:"price_override"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   *string    `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
}

type ViewRoom struct {
	RoomName       string     `json:"room_name" db:"room_name"`
	TypeID         string     `json:"type_id" db:"type_id"`
	Floor          int        `json:"floor" db:"floor"`
	Status         string     `json:"status" db:"status"`
	PriceOverride  float64    `json:"price_override" db:"price_override"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   string     `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
}
type AddRoom struct {
	//RoomID         string `json:"room_id" db:"room_id"`
	RoomName       string     `json:"room_name" db:"room_name"`
	TypeID         string     `json:"type_id" db:"type_id"`
	Floor          int        `json:"floor" db:"floor"`
	Status         string     `json:"status" db:"status"`
	PriceOverride  float64    `json:"price_override" db:"price_override"`
	CleaningStatus string     `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    *time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   *time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   string     `json:"current_guest" db:"current_guest"`
	Note           string     `json:"note" db:"note"`
	CreateBy       string     `json:"createby" db:"createby"`
}

type UpdateRoom struct {
	RoomName       string    `json:"room_name" db:"room_name"`
	TypeID         string    `json:"type_id" db:"type_id"`
	Floor          int       `json:"floor" db:"floor"`
	Status         string    `json:"status" db:"status"`
	PriceOverride  float64   `json:"price_override" db:"price_override"`
	CleaningStatus string    `json:"cleaning_status" db:"cleaning_status"`
	CheckInTime    time.Time `json:"check_in_time" db:"check_in_time"`
	CheckOutTime   time.Time `json:"check_out_time" db:"check_out_time"`
	CurrentGuest   string    `json:"current_guest" db:"current_guest"`
	Note           string    `json:"note" db:"note"`
	UpdateBy       string    `json:"updateby" db:"updateby"`
}
type DeleteRoom struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
