package model

type Room struct {
	Room_id     int    `json:"room_id" db:"room_id"`
	Room_type   string `json:"room_type" db:"room_type"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Room_status bool   `json:"room_status" db:"room_status"`
	Max_guest   int    `json:"max_guest" db:"max_guest"`
	Image_url   string `json:"image_url" db:"image_url"`
}

type RoomUpdate struct {
	Room_type   string `json:"room_type" db:"room_type"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Room_status bool   `json:"room_status" db:"room_status"`
	Max_guest   int    `json:"max_guest" db:"max_guest"`
	Image_url   string `json:"image_url" db:"image_url"`
}