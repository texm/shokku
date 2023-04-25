package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func createQueryContext(route string, params url.Values) (*env.Env, echo.Context) {
	e := env.NewTestingEnvironment()

	uri := fmt.Sprintf("%s/?%s", route, params.Encode())
	req := httptest.NewRequest(http.MethodGet, uri, nil)
	rec := httptest.NewRecorder()

	c := e.Router.NewContext(req, rec)

	e.Router.GET("/ping", e.H(pingRoute))

	return e, c
}

type pingRequest struct {
	Foo string `query:"foo" validate:"alphanum"`
}

func pingRoute(e *env.Env, c echo.Context) error {
	var req pingRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}
	return c.NoContent(http.StatusOK)
}

func TestPingRouteValidatesSuccess(t *testing.T) {
	q := make(url.Values)
	q.Set("foo", "bar")
	e, c := createQueryContext("/ping", q)

	err := pingRoute(e, c)
	assert.NoError(t, err)
}

func TestPingRouteValidatesFail(t *testing.T) {
	q := make(url.Values)
	q.Set("foo", "bar!!")
	e, c := createQueryContext("/ping", q)

	err := pingRoute(e, c)
	assert.Error(t, err)
}
