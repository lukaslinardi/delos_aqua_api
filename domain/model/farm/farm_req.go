package farm

import "time"

type InsertFarm struct {
	FarmName  string    `json:"farm_name"`
	IsDeleted bool      `json:"is_deleted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


