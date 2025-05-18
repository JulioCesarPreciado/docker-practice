package year

// Year represents a vehicle model year used in the public vehicle catalog.
// It includes the unique ID, year name (e.g., 2023), and the associated model ID.
// This structure is typically used for API responses where users select a specific vehicle year.
type Year struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	ModelID int    `json:"model_id" db:"model_id"`
}
