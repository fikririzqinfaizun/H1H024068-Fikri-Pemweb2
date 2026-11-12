package model

import "time"

type TokenBlacklist struct {
	ID         uint      `gorm:"primaryKey"`
	Token      string    `gorm:"type:text;not null"`
	Kadaluarsa time.Time `gorm:"not null;index"`
	DibuatPada time.Time `gorm:"autoCreateTime"`
}
