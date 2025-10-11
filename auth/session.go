package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Session keys for storing user data
const (
	SessionKeyUserID   = "user_id"
	SessionKeyEmail    = "email"
	SessionKeyAvatar   = "avatar"
	SessionKeyUserMeta = "user_metadata"
	// Additional metadata fields
	SessionKeyName     = "name"
	SessionKeyUsername = "username"
	SessionKeyProvider = "provider"
	// Store metadata as JSON string
	SessionKeyMetadataJSON = "metadata_json"
)

// User represents user info stored in session
type User struct {
	ID           string                 `json:"id"`
	Email        string                 `json:"email"`
	Avatar       string                 `json:"avatar"`
	UserMetadata map[string]interface{} `json:"user_metadata"`
	// Additional fields extracted from metadata
	Name     string `json:"name"`
	Username string `json:"username"`
	Provider string `json:"provider"`
}

// GetSessionName returns the configured session name
func GetSessionName() string {
	name := os.Getenv("SESSION_NAME")
	if name == "" {
		name = "social-ops-session"
	}
	return name
}

// StoreUser saves user data in the session
func StoreUser(c echo.Context, user *User) error {
	sess, err := session.Get(GetSessionName(), c)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	// Store basic user data
	sess.Values[SessionKeyUserID] = user.ID
	sess.Values[SessionKeyEmail] = user.Email
	sess.Values[SessionKeyAvatar] = user.Avatar
	sess.Values[SessionKeyName] = user.Name
	sess.Values[SessionKeyUsername] = user.Username
	sess.Values[SessionKeyProvider] = user.Provider

	// Store metadata as JSON string to avoid serialization issues
	// Only store essential metadata to avoid session size limits
	if user.UserMetadata != nil {
		// Create a filtered metadata map with only essential fields
		essentialMetadata := make(map[string]interface{})
		essentialFields := []string{
			"name", "full_name", "given_name", "family_name",
			"email", "email_verified", "picture", "avatar_url",
			"preferred_username", "nickname", "login", "username",
			"location", "company", "bio", "blog", "website",
			"twitter_username", "html_url",
		}

		for _, field := range essentialFields {
			if val, exists := user.UserMetadata[field]; exists {
				essentialMetadata[field] = val
			}
		}

		metadataJSON, err := json.Marshal(essentialMetadata)
		if err != nil {
			c.Logger().Warnf("Failed to marshal user metadata: %v", err)
		} else if len(metadataJSON) > 4000 { // Session cookie size limit check
			c.Logger().Warnf("User metadata too large for session storage (%d bytes)", len(metadataJSON))
		} else {
			sess.Values[SessionKeyMetadataJSON] = string(metadataJSON)
		}
	}

	// Save session with error handling
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}

	return nil
}

// GetUser retrieves user data from session
func GetUser(c echo.Context) (*User, error) {
	sess, err := session.Get(GetSessionName(), c)
	if err != nil {
		return nil, err
	}

	// Check if session has user data
	userID, ok := sess.Values[SessionKeyUserID].(string)
	if !ok || userID == "" {
		return nil, nil // No user in session
	}

	email, _ := sess.Values[SessionKeyEmail].(string)
	avatar, _ := sess.Values[SessionKeyAvatar].(string)
	name, _ := sess.Values[SessionKeyName].(string)
	username, _ := sess.Values[SessionKeyUsername].(string)
	provider, _ := sess.Values[SessionKeyProvider].(string)

	user := &User{
		ID:       userID,
		Email:    email,
		Avatar:   avatar,
		Name:     name,
		Username: username,
		Provider: provider,
	}

	// Retrieve metadata from JSON string
	if metadataJSON, ok := sess.Values[SessionKeyMetadataJSON].(string); ok && metadataJSON != "" {
		var metadata map[string]interface{}
		if err := json.Unmarshal([]byte(metadataJSON), &metadata); err == nil {
			user.UserMetadata = metadata
		}
	}

	return user, nil
}

// ClearSession removes all session data
func ClearSession(c echo.Context) error {
	sess, err := session.Get(GetSessionName(), c)
	if err != nil {
		return err
	}

	// Clear all values
	sess.Values = make(map[interface{}]interface{})
	sess.Options = &sessions.Options{
		MaxAge: -1, // Delete cookie
	}

	return sess.Save(c.Request(), c.Response())
}

// IsAuthenticated checks if user is logged in
func IsAuthenticated(c echo.Context) bool {
	user, err := GetUser(c)
	return err == nil && user != nil
}

// ExtractUserDataFromMetadata extracts common fields from user metadata
// This function handles different OAuth provider metadata formats:
//
// Google provides:
//   - name, full_name: User's full name
//   - given_name: First name
//   - family_name: Last name
//   - picture: Avatar URL
//   - email_verified: Boolean
//
// GitHub provides:
//   - name: Full name (may be empty)
//   - login: Username
//   - avatar_url: Avatar URL
//   - company: Company name (may include @)
//   - location: Location string
//   - bio: User bio
//   - blog: Website URL
//   - twitter_username: Twitter handle (without @)
//   - html_url: GitHub profile URL
//
// Facebook provides:
//   - name: Full name
//   - first_name: First name
//   - last_name: Last name
//   - picture.data.url: Avatar URL (nested)
//   - email: Email address
//
// Twitter/X provides:
//   - name: Display name
//   - username: Twitter handle
//   - profile_image_url: Avatar URL
func ExtractUserDataFromMetadata(metadata map[string]interface{}) (name, username, provider string) {
	// Extract name - try multiple fields
	if n, ok := metadata["name"].(string); ok && n != "" {
		name = n
	} else if n, ok := metadata["full_name"].(string); ok && n != "" {
		name = n
	} else {
		// Try to combine first and last name
		firstName, _ := metadata["given_name"].(string)
		if firstName == "" {
			firstName, _ = metadata["first_name"].(string)
		}
		lastName, _ := metadata["family_name"].(string)
		if lastName == "" {
			lastName, _ = metadata["last_name"].(string)
		}
		if firstName != "" || lastName != "" {
			name = strings.TrimSpace(firstName + " " + lastName)
		}
	}

	// Extract username - try multiple fields
	if u, ok := metadata["preferred_username"].(string); ok && u != "" {
		username = u
	} else if u, ok := metadata["nickname"].(string); ok && u != "" {
		username = u
	} else if u, ok := metadata["login"].(string); ok && u != "" {
		username = u
	} else if u, ok := metadata["username"].(string); ok && u != "" {
		username = u
	}

	// Extract provider - check various possible locations
	if p, ok := metadata["provider"].(string); ok && p != "" {
		provider = p
	} else if p, ok := metadata["iss"].(string); ok {
		// Extract provider from issuer URL
		if strings.Contains(p, "github") {
			provider = "github"
		} else if strings.Contains(p, "google") {
			provider = "google"
		}
	}

	return name, username, provider
}

// GetUserDisplayName returns the best display name for a user
func GetUserDisplayName(user *User) string {
	if user == nil {
		return "Unknown User"
	}

	// Priority: Name → Username → Email → Fallback
	if user.Name != "" {
		return user.Name
	}
	if user.Username != "" {
		return user.Username
	}
	if user.Email != "" {
		return user.Email
	}

	return "User" // Final fallback
}