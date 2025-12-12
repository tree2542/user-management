package domain

// # Enterprise Business Rules (Entities)
type User struct {
	Model
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
}
