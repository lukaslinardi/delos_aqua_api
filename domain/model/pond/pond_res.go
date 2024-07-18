package pond

type Ponds struct {
	ID        int    `json:"id" db:"id"`
	FarmID    int    `json:"farm_id" db:"farm_id"`
	PondName  string `json:"pond_name" db:"pond_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
