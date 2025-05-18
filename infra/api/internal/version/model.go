package version

// Version represents a vehicle version used in the public vehicle catalog.
// It includes the unique ID, version name, and the associated year ID.
// This structure is typically used for API responses where users select a specific vehicle version.
type Version struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	YearID int    `json:"year_id" db:"year_id"`
}
