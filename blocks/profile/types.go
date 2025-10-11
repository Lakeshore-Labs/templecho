package profile

// ProfileConfig contains all configuration for the profile component
type ProfileConfig struct {
	// User authentication data
	UserID    string
	Email     string
	AvatarURL string
	Provider  string // OAuth provider (currently only "google")

	// Personal information from Google OAuth
	FirstName string
	LastName  string

	// Metadata from Supabase
	UserMetadata map[string]interface{}

	// UI Configuration
	ReadOnly   bool   // If true, show as view-only profile
	FormAction string // Where to POST form data
	ShowDebug  bool   // Show raw metadata for debugging
}
