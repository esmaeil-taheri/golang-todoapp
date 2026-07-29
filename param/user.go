// Package param holds the transport-facing request/response types of each use
// case. They carry json tags so the very same structs work for the CLI today
// and the HTTP handlers you'll add next — keeping entity free of transport tags.
package param

type RegisterRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// UserInfo is a safe view of a user: it deliberately omits the password hash.
type UserInfo struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	User UserInfo `json:"user"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User UserInfo `json:"user"`
}
