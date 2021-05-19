package errgroup

import "fmt"

type ErrorGroup interface {
	Append(error)
	Errors() []error
	Error() string
	Empty() bool
}

type errorGroup struct {
	errors []error
	name   string
}

func NewErrorGroup(name string) ErrorGroup {
	return &errorGroup{
		name: name,
	}
}

func (g *errorGroup) Append(err error) {
	if err != nil {
		g.errors = append(g.errors, err)
	}
}

func (g *errorGroup) Empty() bool {
	return len(g.errors) == 0
}

func (g *errorGroup) Errors() []error {
	return g.errors
}

func (g *errorGroup) Error() string {
	return fmt.Sprint(g.name, g.errors)
}
