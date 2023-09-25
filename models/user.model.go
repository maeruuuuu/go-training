package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key;not null;autoIncrement:true"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
