package domain

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Model struct {
	Uid       *uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (m *Model) SetModel() {
	fmt.Println("is Model")
	genUid := uuid.NewV4()
	m.Uid = &genUid
}
