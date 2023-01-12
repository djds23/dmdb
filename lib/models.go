package lib

// Move these to models.go
// This func must be Exported, Capitalized, and comment added.
type Credits struct {
	Cast []Actor `json:"cast"`
	Id   int     `json:"id"`
}

// This func must be Exported, Capitalized, and comment added.
type Actor struct {
	Popularity  float64 `json:"popularity"`
	Name        string  `json:"name"`
	Id          int     `json:"id"`
	ProfilePath string  `json:"profile_path,omitempty"`
}

// This func must be Exported, Capitalized, and comment added.
type Movie struct {
	OriginalTitle string  `json:"original_title"`
	Id            float64 `json:"id"`
}

// This func must be Exported, Capitalized, and comment added.
type Movies struct {
	Results []Movie `json:"results"`
	Page    float64 `json:"page"`
}
