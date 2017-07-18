package log

import (
	"io"

	"github.com/pkg/errors"
)

type filelogger struct {
	*baseLogger
}

func InitFileLogger(w io.WriteCloser, prefix string) (filelogger, error) {
	b, err := InitBaseLogger(w, prefix)
	if err != nil {
		return filelogger{}, errors.Wrapf(err, "error while initializing base logger")
	}

	return filelogger{
		baseLogger: b,
	}, nil
}

func (f filelogger) Close() {
	v, ok := f.baseLogger.w.(io.WriteCloser)
	if ok {
		v.Close()
	}
}
