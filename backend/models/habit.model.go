package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID        uuid.UUID        `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID        `json:"user_id" gorm:"not null"`
	Name      *string          `json:"name" gorm:"not null;size:100"`
	Goal      *string          `json:"goal" gorm:"not null;size:255"`
	Time      *json.RawMessage `json:"time" gorm:"not null;type:jsonb"`
	CreatedAt time.Time        `json:"created_at"`
}
