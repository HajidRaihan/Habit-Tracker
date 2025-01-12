package requests

import (
	"encoding/json"
)

type CreateHabitRequest struct {
	Name      string          `json:"name" binding:"required"`
	Goal      string          `json:"goal" binding:"required"`
	Time      json.RawMessage `json:"time" binding:"required"` // Menyimpan data waktu sebagai raw JSON
	CreatedAt string          `json:"created_at"`
}
