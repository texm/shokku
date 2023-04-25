package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const DefaultBCryptCost = 14

type PasswordAuthenticator struct {
	baseAuthenticator
	bcryptCost int
}

func NewPasswordAuthenticator(cfg Config, bCryptCost int) (*PasswordAuthenticator, error) {
	pwAuth := &PasswordAuthenticator{}
	pwAuth.bcryptCost = bCryptCost
	pwAuth.signingKey = cfg.SigningKey
	pwAuth.tokenLifetime = cfg.TokenLifetime
	pwAuth.cookieDomain = cfg.CookieDomain
	pwAuth.authMethod = MethodPassword

	return pwAuth, nil
}

func (a *PasswordAuthenticator) HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, a.bcryptCost)
}

func (a *PasswordAuthenticator) VerifyHash(password []byte, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, password) == nil
}
