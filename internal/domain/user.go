package domain

// # Enterprise Business Rules (Entities)
type User struct {
	Model
	FirstName string
	LastName  string
	Username  string `gorm:"uniqueIndex"`
	Password  string
	Email     string `gorm:"uniqueIndex"`
	Role      string
}
