package dto

type Material struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
