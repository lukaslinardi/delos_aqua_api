package farm

import (
	"github.com/lukaslinardi/delos_aqua_api/domain/model/pond"
)

type Farms struct {
	ID        int    `json:"id" db:"id"`
	FarmName  string `json:"farm_name" db:"farm_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Farm struct {
	ID        int    `json:"id" db:"id"`
	PondID    int    `json:"pond_id" db:"pond_id"`
	FarmName  string `json:"farm_name" db:"farm_name"`
	PondName  string `json:"pond_name" db:"pond_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type FarmRes struct {
	ID       int          `json:"id" db:"id"`
	FarmName string       `json:"farm_name" db:"farm_name"`
	Ponds    []pond.Ponds `json:"ponds"`
}
