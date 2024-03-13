package models

import "time"

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"senha,omitempty"`
	CreatedA time.Time `json:"data_create,omitempty"`
}
