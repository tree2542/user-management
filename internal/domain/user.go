package domain

import "time"

//# Enterprise Business Rules (Entities)
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
