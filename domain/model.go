package domain

type Gifter interface {
	Name() string
}

type Group interface {
	ID() string
	Name() string
	Budget() float64
	Gifters() []Gifter
	AddGifter(Gifter) error
}

type gifter struct {
	name string
}

func NewGifter(name string) Gifter {
	return &gifter{
		name: name,
	}
}

func (g *gifter) Name() string {
	return g.name
}

type group struct {
	id      string
	name    string
	budget  float64
	gifters []Gifter
}

func NewGroup(id, name string, budget float64) Group {
	return &group{
		id:     id,
		name:   name,
		budget: budget,
	}
}

func (grp *group) ID() string {
	return grp.id
}
func (grp *group) Name() string {
	return grp.name
}
func (grp *group) Budget() float64 {
	return grp.budget
}
func (grp *group) Gifters() []Gifter {
	return grp.gifters
}
func (grp *group) AddGifter(g Gifter) error {
	grp.gifters = append(grp.gifters, g)
	return nil
}
