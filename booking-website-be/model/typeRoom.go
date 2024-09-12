package model

type TypeRoom struct {
	//TypeID        string    `json:"type_id" db:"type_id"`
	TypeName      string  `json:"type_name" db:"type_name"`
	Description   string  `json:"description" db:"description"`
	PricePerNight float32 `json:"price_per_night" db:"price_per_night"`
	MaxOccupancy  int     `json:"max_occupancy" db:"max_occupancy"`
	RoomSize      float32 `json:"room_size" db:"room_size"`
	ImageURL      string  `json:"image_url" db:"image_url"`
	Status        string  `json:"status" db:"status"`
	Discount      float64 `json:"discount" db:"discount"`
	CreateBy      string  `json:"createby" db:"createby"`
}

type SelectTypeRoom struct {
	TypeID        string  `json:"type_id" db:"type_id"`
	TypeName      string  `json:"type_name" db:"type_name"`
	Description   string  `json:"description" db:"description"`
	PricePerNight float32 `json:"price_per_night" db:"price_per_night"`
	MaxOccupancy  int     `json:"max_occupancy" db:"max_occupancy"`
	RoomSize      float32 `json:"room_size" db:"room_size"`
	ImageURL      string  `json:"image_url" db:"image_url"`
	Status        string  `json:"status" db:"status"`
	Discount      float64 `json:"discount" db:"discount"`
}

type UpdateTypeRoom struct {
	TypeName      string  `json:"type_name" db:"type_name"`
	Description   string  `json:"description" db:"description"`
	PricePerNight float32 `json:"price_per_night" db:"price_per_night"`
	MaxOccupancy  int     `json:"max_occupancy" db:"max_occupancy"`
	RoomSize      float32 `json:"room_size" db:"room_size"`
	ImageURL      string  `json:"image_url" db:"image_url"`
	Status        string  `json:"status" db:"status"`
	Discount      float64 `json:"discount" db:"discount"`
	UpdateBy      string  `json:"updateby" db:"updateby"`
}

type DeleteTypeRoom struct {
	DeleteBy string `json:"deleteby" db:"deleteby"`
}
