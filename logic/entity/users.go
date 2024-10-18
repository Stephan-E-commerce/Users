package entity

type User struct {
	ID           int    `json:"id" db:"id"`
	Email        string `json:"email" db:"email"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}

type UserGRPC struct {
	Name         string
	Email        string
	PasswordHash string
}
