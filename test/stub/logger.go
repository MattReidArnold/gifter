package stub

import "github.com/mattreidarnold/gifter/app"

type stubLogger struct{}

func NewStubLogger() app.Logger {
	return &stubLogger{}
}

func (l *stubLogger) Info(args ...interface{})             {}
func (l *stubLogger) Error(err error, args ...interface{}) {}
