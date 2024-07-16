package farm

type Farms struct {
	ID        int    `json:"id" db:"id"`
	FarmName  string `json:"farm_name" db:"farm_name"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
