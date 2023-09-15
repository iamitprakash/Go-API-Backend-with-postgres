package models

type Claims struct {
	Username  string `json:"username"`
	ExpiresAt int64  `json:"expires_at"`
	Role      string `json:"roles`
}
