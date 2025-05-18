package model

// Model represents a vehicle model used in the public vehicle catalog.
// It includes the unique ID, displayable name, and the associated brand ID.
// This structure is typically used for API responses where model selection is needed.
type Model struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	BrandID int    `json:"brand_id" db:"brand_id"`
}
