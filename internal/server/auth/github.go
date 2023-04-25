package auth

type GithubAuthenticator struct {
	baseAuthenticator
}

func NewGithubAuthenticator(cfg Config) (*GithubAuthenticator, error) {
	ghAuth := &GithubAuthenticator{}
	// TODO: check these
	ghAuth.signingKey = cfg.SigningKey
	ghAuth.tokenLifetime = cfg.TokenLifetime
	ghAuth.cookieDomain = cfg.CookieDomain
	ghAuth.authMethod = MethodGithub

	return ghAuth, nil
}
