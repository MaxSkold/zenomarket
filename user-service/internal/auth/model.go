package auth

type Account struct {
	ID           int64   `json:"id" db:"id"`
	UserID       int64   `json:"user_id" db:"user_id"`
	PasswordHash string  `json:"-" db:"password_hash"`
	Provider     string  `json:"provider" db:"provider"`
	ThirdPartyID *string `json:"third_party_id,omitempty" db:"third_party_id"`
	LastLoginAt  *string `json:"last_login_at,omitempty" db:"last_login_at"`
	UpdatedAt    string  `json:"updated_at" db:"updated_at"`
}
