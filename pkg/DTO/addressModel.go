package dto

// Address represents the table of address in the database schema.
type Address struct {
	House  string `json:"house"`
	Street string `json:"street"`
	City   string `json:"city"`
	ZIP    uint   `json:"zip"`
	State  string `json:"state"`
	// UserID uint   `json:"user_id"`
}
