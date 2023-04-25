package auth

import (
	"net/http"
	"time"
)

type authCookieValues struct {
	name  string
	value string
	// domain   string
	httpOnly bool
	lifetime time.Duration
}

func makeAuthCookie(values authCookieValues) *http.Cookie {
	expiresAt := time.Now().Add(values.lifetime)

	cookie := &http.Cookie{
		Name:     values.name,
		Value:    values.value,
		Expires:  expiresAt,
		HttpOnly: values.httpOnly,
		Path:     "/",
		// Domain:   values.domain,
		MaxAge:   int(values.lifetime.Seconds()),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}

func clearAuthCookie(name string, httpOnly bool, domain string) *http.Cookie {
	vals := authCookieValues{
		name:  name,
		value: "",
		// domain:   domain,
		httpOnly: httpOnly,
		lifetime: time.Second,
	}
	cookie := makeAuthCookie(vals)
	return cookie
}
