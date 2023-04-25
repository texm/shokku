package commands

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/server/dto"
	"io"
	"regexp"
	"strings"
	"time"
	"unicode"
)

type AsyncDokkuCommand func() (*dokku.CommandOutputStream, error)

const ansiRegexP = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	executions  = map[string]execution{}
	statuses    = map[string]*ExecutionStatus{}
	ansiRe      = regexp.MustCompile(ansiRegexP)
	fiveMinutes = time.Minute * 5
	readTimeout = time.Second * 5

	bufSize = 4096
)

var (
	ErrNotPolled   = errors.New("not polled yet")
	ErrNoExecution = errors.New("no such command id")
)

func sanitiseOutput(b []byte) []byte {
	b = ansiRe.ReplaceAll(b, []byte{})
	mapped := bytes.Map(func(r rune) rune {
		if r > unicode.MaxASCII {
			return -1
		}
		return r
	}, b)
	return mapped
}

func readOutput(s chan string, e chan error, r io.Reader) {
	// defer close(s)
	// defer close(e)

	output := bytes.Buffer{}
	reader := bufio.NewReader(r)
	for {
		buf := make([]byte, bufSize)
		n, err := reader.Read(buf)
		if err != nil {
			e <- err
			break
		}
		if n > 0 {
			output.Write(sanitiseOutput(buf[:n]))
		}
		if n < bufSize {
			break
		}
	}
	if output.Len() > 0 {
		s <- strings.TrimSpace(output.String())
	}
}

func ReadWithTimeout(reader io.Reader, timeout time.Duration) (string, error) {
	s := make(chan string)
	e := make(chan error)

	go readOutput(s, e, reader)

	select {
	case str := <-s:
		return str, nil
	case err := <-e:
		return "", err
	case <-time.After(timeout):
		return "", nil
	}
}

func toOutputLines(lines []string, outputType string, pollTime time.Time) []dto.OutputLine {
	output := make([]dto.OutputLine, len(lines))
	for i, line := range lines {
		output[i] = dto.OutputLine{Msg: line, Type: outputType, PolledAt: pollTime}
	}
	return output
}

func getExecutionStatus(id string) *ExecutionStatus {
	status, ok := statuses[id]
	if !ok {
		return &ExecutionStatus{}
	}
	return status
}

func PollStatuses() {
	for id, status := range statuses {
		if !status.Finished {
			continue
		}
		if time.Since(status.FinishedAt) > fiveMinutes {
			delete(statuses, id)
		}
	}

	for id, exec := range executions {
		if exec.output == nil {
			continue
		}
		status := getExecutionStatus(id)
		pollTime := time.Now()

		stdout, stdoutErr := ReadWithTimeout(exec.output.Stdout, readTimeout)
		if stdoutErr != nil {
			status.StdoutReadError = stdoutErr
		}
		if stdout != "" {
			stdoutLines := strings.Split(stdout, "\n")
			output := toOutputLines(stdoutLines, "stdout", pollTime)

			// status.Stdout = append(status.Stdout, stdoutLines...)
			status.CombinedOutput = append(status.CombinedOutput, output...)
		}

		stderr, stderrErr := ReadWithTimeout(exec.output.Stderr, readTimeout)
		if stderrErr != nil {
			status.StderrReadError = stderrErr
		}
		if stderr != "" {
			stderrLines := strings.Split(stderr, "\n")
			// status.Stderr = append(status.Stderr, stderrLines...)
			output := toOutputLines(stderrLines, "stderr", pollTime)
			status.CombinedOutput = append(status.CombinedOutput, output...)
		}

		status.StreamError = exec.output.Error

		if errors.Is(stdoutErr, io.EOF) {
			delete(executions, id)
			status.Finished = true
			status.FinishedAt = time.Now()
			if exec.callback != nil {
				status.CallbackError = exec.callback()
			}
		}
		statuses[id] = status
	}

	time.Sleep(time.Second)
	PollStatuses()
}
