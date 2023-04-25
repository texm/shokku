package auth

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

const (
	ContextUserKey      = "user"
	DataCookieName      = "auth_data"
	SignatureCookieName = "auth_sig"
)

type Method string

const (
	MethodNone     = Method("none")
	MethodPassword = Method("password")
	MethodGithub   = Method("github")
)

var (
	ErrNoTokenInContext = errors.New("no token value in context")
	ErrNoUserInContext  = errors.New("failed to retrieve user from context")
	ErrClaimsInvalid    = errors.New("failed to cast jwt claims")
)

type Authenticator interface {
	NewToken(claims UserClaims) (string, error)
	SetUserContext(c echo.Context, contextKey string) error
	GetUserFromContext(c echo.Context) (*User, error)
	SetTokenCookies(c echo.Context, jwt string) string
	ClearTokenCookies(c echo.Context)
	GetSigningKey() []byte
	GetCookieDomain() string
	GetTokenLifetime() time.Duration
	GetMethod() Method
}

type User struct {
	UserClaims
	jwt.StandardClaims
}

type UserClaims struct {
	Name string `json:"name"`
}

type baseAuthenticator struct {
	signingKey    []byte
	authMethod    Method
	cookieDomain  string
	tokenLifetime time.Duration
}

type Config struct {
	SigningKey    []byte
	CookieDomain  string
	TokenLifetime time.Duration
}

func (a *baseAuthenticator) NewToken(claims UserClaims) (string, error) {
	expiry := time.Now().Add(a.tokenLifetime)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expiry.Unix(),
	}

	user := &User{
		UserClaims:     claims,
		StandardClaims: stdClaims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString(a.signingKey)
}

func (a *baseAuthenticator) SetUserContext(c echo.Context, contextKey string) error {
	token, ok := c.Get(contextKey).(*jwt.Token)
	if !ok {
		return ErrNoTokenInContext
	}

	u, ok := token.Claims.(*User)
	if !ok {
		return ErrClaimsInvalid
	}

	c.Set(ContextUserKey, u)

	return nil
}

func (a *baseAuthenticator) GetUserFromContext(c echo.Context) (*User, error) {
	user, ok := c.Get(ContextUserKey).(*User)
	if !ok {
		log.Error().Msg("failed to retrieve user from context")
		return nil, ErrNoUserInContext
	}
	return user, nil
}

func (a *baseAuthenticator) SetTokenCookies(c echo.Context, jwt string) string {
	splitToken := strings.Split(jwt, ".")
	dataCookieValue := strings.Join(splitToken[:2], ".")
	signatureCookieValue := splitToken[2]

	// accessible to the js frontend
	dataCookieValues := authCookieValues{
		name:     DataCookieName,
		value:    dataCookieValue,
		httpOnly: false,
		lifetime: a.tokenLifetime,
	}
	c.SetCookie(makeAuthCookie(dataCookieValues))

	// inaccessible to the js frontend
	signatureCookieValues := authCookieValues{
		name:     SignatureCookieName,
		value:    signatureCookieValue,
		httpOnly: true,
		lifetime: a.tokenLifetime,
	}
	c.SetCookie(makeAuthCookie(signatureCookieValues))

	return dataCookieValue
}

func (a *baseAuthenticator) ClearTokenCookies(c echo.Context) {
	c.SetCookie(clearAuthCookie(DataCookieName, false, a.cookieDomain))
	c.SetCookie(clearAuthCookie(SignatureCookieName, true, a.cookieDomain))
}

func (a *baseAuthenticator) GetSigningKey() []byte {
	return a.signingKey
}

func (a *baseAuthenticator) GetMethod() Method {
	return a.authMethod
}

func (a *baseAuthenticator) GetCookieDomain() string {
	return a.cookieDomain
}

func (a *baseAuthenticator) GetTokenLifetime() time.Duration {
	return a.tokenLifetime
}
