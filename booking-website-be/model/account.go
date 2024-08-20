package model

type SaveAccount struct {
	Name     string `json:"name" db:"name"`
	Dob      string `json:"dob" db:"dob"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
	Role     string `json:"role" db:"role"`
}

type SignIn struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}
