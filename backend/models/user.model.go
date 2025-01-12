package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
}
