package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Gender    string `gorm:"column:gender"`
	Phone     string `gorm:"column:phone"`
	Email     string `gorm:"column:email"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UID       sql.NullString `json:"user_id" swaggertype:"string"`
}
