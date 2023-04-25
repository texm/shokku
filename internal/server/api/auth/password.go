package auth

import (
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	auth "gitlab.com/texm/shokku/internal/server/auth"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/models"
)

func HandlePasswordLogin(e *env.Env, c echo.Context) error {
	var req dto.PasswordLoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	var dbUser models.User
	res := e.DB.Where("name = ?", req.Username).Take(&dbUser)
	if res.Error != nil {
		return echo.ErrForbidden
	}

	pwAuth, ok := e.Auth.(*auth.PasswordAuthenticator)
	if !ok {
		log.Error().Msg("failed to cast authenticator to pw auth")
		return echo.ErrInternalServerError
	}
	if !pwAuth.VerifyHash([]byte(req.Password), dbUser.PasswordHash) {
		return echo.ErrForbidden
	}

	if dbUser.TotpEnabled {
		if !totp.Validate(req.TotpCode, dbUser.TotpSecret) {
			return c.JSON(http.StatusOK, dto.PasswordLoginResponse{
				Success:   true,
				NeedsTotp: true,
			})
		}
	}

	claims := auth.UserClaims{
		Name: req.Username,
	}
	token, err := e.Auth.NewToken(claims)
	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt")
		return echo.ErrInternalServerError
	}
	e.Auth.SetTokenCookies(c, token)

	return c.JSON(http.StatusOK, dto.PasswordLoginResponse{
		Success: true,
	})
}
