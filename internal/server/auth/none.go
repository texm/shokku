package auth

type NoneAuthenticator struct {
	baseAuthenticator
}

func NewNoneAuthenticator(cfg Config) (*NoneAuthenticator, error) {
	noneAuth := &NoneAuthenticator{}
	noneAuth.signingKey = cfg.SigningKey
	noneAuth.tokenLifetime = cfg.TokenLifetime
	noneAuth.cookieDomain = cfg.CookieDomain
	noneAuth.authMethod = MethodNone
	return noneAuth, nil
}
