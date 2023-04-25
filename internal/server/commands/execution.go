package commands

import (
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/server/dto"
	"math/rand"
	"time"
)

type CallbackFunc func() error

type execution struct {
	Id       string
	output   *dokku.CommandOutputStream
	callback CallbackFunc
	error    error
}

type ExecutionStatus struct {
	CombinedOutput []dto.OutputLine

	Finished   bool
	FinishedAt time.Time

	// TODO: this smells
	CallbackError   error
	StreamError     error
	StdoutReadError error
	StderrReadError error
}

func generateCommandExecutionId() string {
	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RequestExecution(cmd AsyncDokkuCommand, callback CallbackFunc) string {
	id := generateCommandExecutionId()

	exec := execution{
		Id:       id,
		callback: callback,
	}
	executions[id] = exec

	go func(exec execution, id string) {
		stream, err := cmd()
		exec.output = stream
		if err != nil {
			exec.error = err
		}
		executions[id] = exec
	}(exec, id)

	return id
}

func GetExecutionStatus(id string) (*ExecutionStatus, error) {
	status, ok := statuses[id]
	if !ok {
		_, exists := executions[id]
		if exists {
			return nil, ErrNotPolled
		}
		return nil, ErrNoExecution
	}

	return status, nil
}
