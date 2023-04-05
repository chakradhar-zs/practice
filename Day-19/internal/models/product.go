package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Quantity    int    `json:"q"`
	Category    string `json:"category"`
	Brand       Brand  `json:"brand"`
	Status      string `json:"status"`
}
