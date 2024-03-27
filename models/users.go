package models

import "time"

type User struct {
	UserID        int       `json:"user_id,omitempty"`
	UserFullName  string    `json:"user_full_name,omitempty"`
	UserEmail     string    `json:"user_email,omitempty"`
	UserPassword  string    `json:"user_password,omitempty"`
	UserCreatedAt time.Time `json:"user_created_at,omitempty"`
	UserUpdatedAt time.Time `json:"user_updated_at,omitempty"`
}

type Users []*User
