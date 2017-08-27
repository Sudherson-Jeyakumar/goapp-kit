package os

import (
	"context"
	"io"
	"os/exec"

	"github.com/pkg/errors"
)

type OSCommandFunc func(context.Context, io.Writer, io.Writer) error

type CommandExecutor interface {
	Execute(r io.Reader) (OSCommandFunc, error)
}

type simpleExecutor struct {
	command          string
	args             []string
	isFailedContinue bool
}

func getSimpleExecutor() *simpleExecutor {
	return &simpleExecutor{
		isFailedContinue: false,
	}
}

func (s *simpleExecutor) IfFailedContinue(f bool) *simpleExecutor {
	s.isFailedContinue = f
	return s
}

func (s *simpleExecutor) Execute(r io.Reader) (OSCommandFunc, error) {
	dummyFunc := func(context.Context, io.Writer, io.Writer) error { return nil }

	if s.command == "" {
		return dummyFunc, errors.Errorf("Command Not specified;")
	}

	return func(ctx context.Context, outw io.Writer, errw io.Writer) error {

		c := exec.CommandContext(ctx, s.command, s.args...)

		sout, err := c.StdoutPipe()
		if err != nil {
			return errors.Wrapf(err, "error while setting Output Writer")
		}

		serr, err := c.StderrPipe()
		if err != nil {
			return errors.Wrapf(err, "error while setting Error Writer")
		}

		stdin, err := c.StdinPipe()
		if err != nil {
			return errors.Wrapf(err, "error while getting Stdin Pipe")
		}

		go func(r io.Reader) {
			defer stdin.Close()

			if r != nil {
				io.Copy(stdin, r)
			}
		}(r)

		if err = c.Start(); err != nil {
			return errors.Wrapf(err, "error start executing command")
		}

		if _, err = io.Copy(outw, sout); err != nil {
			return errors.Wrapf(err, "error while copying output to writer")
		}

		if _, err = io.Copy(errw, serr); err != nil {
			return errors.Wrapf(err, "error while copying error to writer")
		}

		if err = c.Wait(); err != nil {
			return errors.Wrapf(err, "error while executing command")
		}

		return nil

	}, nil
}
