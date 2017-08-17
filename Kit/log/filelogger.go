package log

import (
	"fmt"
	"log"
	"os"
	"strings"

	kitcfg "Kit/cfg"

	"github.com/pkg/errors"
)

type filelogger struct {
	*baseLogger
	close func() error
}

func (f *filelogger) Close() {
	f.close()
}

func WithFile() *filelogger {
	return &filelogger{
		baseLogger: baseLoggerObj,
		close:      func() error { return nil },
	}
}

func initLoggerWithFile(fileName string) (*filelogger, error) {
	l := WithFile()

	err := l.apply(fileName)
	if err != nil {
		return nil, errors.Wrapf(err, "error initializing logger;")
	}

	return l, nil
}

func (f *filelogger) apply(v string) error {

	if len(strings.Trim(v, " ")) == 0 {
		return errors.Errorf("Logger FileName cannot be empty")
	}

	fl, err := os.OpenFile(v, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return errors.Wrapf(err, "error while initializing file logger")
	}

	fmt.Println("Initializing fileLogger", *(f.baseLogger))

	f.baseLogger.w = fl
	f.baseLogger.logger = log.New(f.baseLogger.w, "", log.LstdFlags)
	f.close = func() error {
		return fl.Close()
	}

	return nil
}

func (f *filelogger) Apply(c kitcfg.Config) error {
	writerType := c.Value("loggers", "file")

	if writerType == nil {
		return errors.Errorf("config setting not found; for file logger")
	}

	if v, ok := writerType.(string); ok {
		f.apply(v)
	} else {
		return errors.Errorf("error while initializing file logger - filename not present")
	}

	return nil
}
