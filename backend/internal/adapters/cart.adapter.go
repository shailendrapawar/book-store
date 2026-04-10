package adapters

import "time"

type Cart struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartSearchFilters struct {
	UserID string `json:"user_id"`
	Status string `json:"status"`
}
