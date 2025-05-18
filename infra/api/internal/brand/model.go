package brand

// Brand represents a vehicle brand used in the public vehicle catalog.
// It includes the unique ID and the displayable name of the brand,
// and is intended to be used in frontend dropdowns or selection lists.
type Brand struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
