package auth

type TokenStore interface {
	GetClientID() string
	GetClientSecret() string
	GetRefreshToken() string
	RotateRefreshToken(string) error
}
