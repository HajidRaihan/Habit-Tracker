package requests

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time

const layout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(ct).Format(layout))), nil
}

// HabitID   uuid.UUID  `json:"habit_id"`
type HabitLogRequest struct {
	LogDate   CustomTime `json:"log_date" binding:"required"`
	Progress  int        `json:"progress" binding:"required"`
	Status    string     `json:"status" enum:"in progress,completed,missed,skipped,partially completed" binding:"required"`
	CreatedAt string     `json:"created_at" `
}
