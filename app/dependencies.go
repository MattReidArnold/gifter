package app

type Dependencies struct {
	Logger
	GroupRepository
}

type Logger interface {
	Info(s string)
}
