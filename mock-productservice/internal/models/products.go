package models

type Product struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	BrandId int    `json:"brand_id"`
}
