package os

import (
	"context"
	"io"

	"github.com/pkg/errors"
)

func SeriallyExecute(ctx context.Context, outw io.ReadWriter, errw io.Writer, ces ...CommandExecutor) error {

	for _, ce := range ces {
		f, err := ce.Execute(nil)
		if err != nil {
			return errors.Wrapf(err, "error while executing commands serially")
		}

		err = f(ctx, outw, errw)
		if err != nil {
			return errors.Wrapf(err, "error while executing commands serially")
		}
	}

	return nil
}

func PipeExecute(ctx context.Context, outw io.ReadWriter, errw io.Writer, ces ...CommandExecutor) error {
	var tempReader io.Reader

	for _, ce := range ces {
		f, err := ce.Execute(tempReader)
		if err != nil {
			return errors.Wrapf(err, "error while executing commands using pipe")
		}

		err = f(ctx, outw, errw)
		if err != nil {
			return errors.Wrapf(err, "error while executing commands using pipe")
		}

		tempReader = outw
	}

	return nil
}
