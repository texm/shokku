package server

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

type wsErrorFunc func(w http.ResponseWriter, r *http.Request, status int, reason error)

func createWSCheckOriginFunc(origin string) func(r *http.Request) bool {
	return func(req *http.Request) bool {
		log.Debug().
			Str("origin", req.Header.Get("Origin")).
			Str("host", origin).
			Msg("ws check origin")
		return true
	}
}

func createWSErrorFunc(e *echo.Echo) wsErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		err := echo.NewHTTPError(status, reason)
		c := e.NewContext(r, w)
		e.HTTPErrorHandler(err, c)
	}
}

func initWebsocketUpgrader(cfg Config, router *echo.Echo) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: createWSCheckOriginFunc(cfg.Host),
		Error:       createWSErrorFunc(router),
	}
}
