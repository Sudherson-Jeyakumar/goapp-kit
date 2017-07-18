package log

import (
	"io"
	"log"
	"strings"

	"github.com/pkg/errors"
)

type Logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
	Tracef(string, ...interface{})
}

type Prefixer interface {
	Prefix(string) error
	PrefixLogType(string, string) error
}

type Enabler interface {
	Enable(bool)
}

type Closer interface {
	Close() error
}

type baseLogger struct {
	logger *log.Logger
	w      io.Writer

	enabled  bool
	prefixes map[string]string
}

func InitBaseLogger(w io.Writer, prefix string) (*baseLogger, error) {
	if w == nil {
		return &baseLogger{}, errors.Errorf("Logger writer cannot be nil")
	}

	b := &baseLogger{
		w:       w,
		logger:  log.New(w, "", log.LstdFlags),
		enabled: true,
		prefixes: map[string]string{
			"info":  "INFO <" + prefix + ">",
			"error": "ERROR <" + prefix + ">",
		},
	}

	return b, nil
}

func (s *baseLogger) Infof(fmt string, args ...interface{}) {
	if s.enabled {
		s.logger.Printf(s.prefixes["info"]+" : "+fmt, args...)
	}
}

func (s *baseLogger) Errorf(fmt string, args ...interface{}) {
	if s.enabled {
		s.logger.Printf(s.prefixes["error"]+" : "+fmt, args...)
	}
}

func (s *baseLogger) Prefix(prefix string) error {
	s.prefixes["info"] = "INFO <" + prefix + ">"
	s.prefixes["error"] = "ERROR <" + prefix + ">"

	return nil
}

func (s *baseLogger) PrefixLogType(logtype string, prefix string) error {
	s.prefixes[logtype] = strings.Split(s.prefixes[logtype], "<")[0] + "<" + prefix + ">"
	return nil
}

func (s *baseLogger) Enabler(f bool) {
	s.enabled = f
}
