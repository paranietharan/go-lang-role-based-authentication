package models

import "time"

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name" validate:"required, min=3, max=30"`
	LastName     string    `json:"last_name" validate:"required, min=3, max=30"`
	Password     string    `json:"password" validate:"required, min=8, max=30"`
	Email        string    `json:"email" validate:"required"`
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	UserType     string    `json:"usertype" validate:"required, eq=USER|eq=ADMIN"`
	UserID       string    `json:"user_id"`
}
