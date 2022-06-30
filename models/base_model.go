package models

import (
	"time"
)

type BaseModel struct {
	ID        string    `json:"id" gorm:"default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
