package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Location  string    `json:"location"`
	Age       int32     `json:"age"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Role      string    `json:"role"`
}
