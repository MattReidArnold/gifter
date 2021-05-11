package log

import (
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/mattreidarnold/gifter/app"
)

type logger struct {
	logger kitlog.Logger
}

func NewLogger(l kitlog.Logger) app.Logger {
	return &logger{
		logger: l,
	}
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Log("msg", fmt.Sprint(args...))
}
