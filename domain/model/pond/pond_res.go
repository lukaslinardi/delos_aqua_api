package pond

type Ponds struct {
	ID        int    `json:"id" db:"id"`
	FarmID    int    `json:"farm_id" db:"farm_id"`
	PondName  string `json:"pond_name" db:"pond_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type Pond struct {
	ID        int    `json:"id" db:"id"`
	FarmID    int    `json:"farm_id" db:"farm_id"`
	FarmName  string `json:"farm_name" db:"farm_name"`
	PondName  string `json:"pond_name" db:"pond_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
