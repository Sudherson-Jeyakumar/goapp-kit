package log

import (
	"Kit/cfg"

	"github.com/pkg/errors"
)

func prefixLogger(l Logger, c cfg.Config, logType string) error {
	if prefixInfo := c.Value("loggers", "prefix", logType); prefixInfo != nil {
		if v, ok := l.(Prefixer); ok {

			p, ok1 := prefixInfo.(string)
			if !ok1 {
				return errors.Errorf("error occurred during prefixing")
			}
			v.PrefixLogType(logType, p)

		} else {
			return errors.Errorf("logger cannot be prefixed")
		}
	}

	return nil
}

func prefixAllLogger(l Logger, c cfg.Config) error {
	if err := prefixLogger(l, c, "info"); err != nil {
		return errors.Wrapf(err, "error while deriving info logger")
	}

	if err := prefixLogger(l, c, "error"); err != nil {
		return errors.Wrapf(err, "error while deriving error logger")
	}
	return nil
}
