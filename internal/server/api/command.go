package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/server/commands"
	"gitlab.com/texm/shokku/internal/server/dto"
	"net/http"
)

func GetCommandExecutionStatus(e *env.Env, c echo.Context) error {
	var req dto.GetCommandExecutionStatusRequest
	if err := dto.BindRequest(c, &req); err != nil {
		return err.ToHTTP()
	}

	status, err := commands.GetExecutionStatus(req.ExecutionID)
	if err != nil {
		if errors.Is(err, commands.ErrNotPolled) {
			return c.JSON(http.StatusOK, dto.CommandExecutionStatusResponse{
				Started: false,
			})
		}
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, dto.CommandExecutionStatusResponse{
		CombinedOutput: status.CombinedOutput,
		Started:        true,
		Finished:       status.Finished,
		Success:        status.StreamError == nil,
	})
}
