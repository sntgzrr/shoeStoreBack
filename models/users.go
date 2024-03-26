package models

import "time"

type User struct {
	UserID        int       `json:"user_id"`
	UserFullName  string    `json:"user_full_name"`
	UserEmail     string    `json:"user_email"`
	UserPassword  string    `json:"user_password"`
	UserCreatedAt time.Time `json:"user_created_at"`
	UserUpdatedAt time.Time `json:"user_updated_at"`
}

type Users []*User
