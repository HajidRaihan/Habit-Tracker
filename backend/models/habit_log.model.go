package models

import (
	"time"

	"github.com/google/uuid"
)

type HabitLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HabitID   uuid.UUID `gorm:"type:uuid" json:"habit_id"`
	LogDate   time.Time `gorm:"type:date;not null" json:"log_date"`
	Progress  int       `gorm:"no null" json:"progress"`
	Status    string    `gorm:"not null" json:"status" enum:"in progress,completed,missed,skipped,partially completed"` //in progress,completed,missed,skipped,partially completed
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
}
