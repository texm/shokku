package dto

import (
	"time"
)

type GetCommandExecutionStatusRequest struct {
	ExecutionID string `query:"execution_id" validate:"required"`
}

type CommandExecutionResponse struct {
	ExecutionID string `json:"execution_id"`
}

type OutputLine struct {
	Msg      string    `json:"message"`
	Type     string    `json:"type"`
	PolledAt time.Time `json:"polled_at"`
}

type CommandExecutionStatusResponse struct {
	Started  bool `json:"started"`
	Finished bool `json:"finished"`
	Success  bool `json:"success"`

	CombinedOutput []OutputLine `json:"output"`
}

type AppExecInProcessRequest struct {
	AppName     string `json:"appName" validate:"appName"`
	ProcessName string `json:"processName" validate:"required"`
	Command     string `json:"command"`
}

type AppExecInProcessResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}
