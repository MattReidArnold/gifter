package app

type Dependencies struct {
	Logger
}

type Logger interface {
	Info(s string)
}
