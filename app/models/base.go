package models

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel defines the basic attributes that every other entity needs to implement
type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key; not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
