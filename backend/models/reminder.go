package models

import (
	"time"

	"github.com/google/uuid"
)

type Reminder struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	HabitID      uuid.UUID `gorm:"type:uuid" json:"habit_id"`
	ReminderTime string    `gorm:"type:time;not null" json:"reminder_time"`
	CreatedAt    time.Time `gorm:"default:now()" json:"created_at"`
}
