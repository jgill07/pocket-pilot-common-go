package auth

import "context"

// PocketUser represents a user in the pocket pilot system
// This is the user that is authenticated to the system
type PocketUser struct {
	UserId      string   `json:"https://pocket-pilot.com/user_id"`
	Email       string   `json:"https://pocket-pilot.com/email"`
	Role        []string `json:"https://pocket-pilot.com/role"`
	Permissions []string `json:"permissions"`
	Subject     string   `json:"sub"`
}

func (p *PocketUser) Validate(_ context.Context) error {
	return nil
}

func (p *PocketUser) HasPermission(permission string) bool {
	for _, perm := range p.Permissions {
		if perm == permission {
			return true
		}
	}
	return false
}
