package pond

import "time"

type InsertPond struct {
	PondName  string    `json:"pond_name"`
	IsDeleted bool      `json:"is_deleted"`
	FarmID    int       `json:"farm_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
