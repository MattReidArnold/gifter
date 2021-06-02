package domain

type Gifter interface {
	ID() string
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
	id   string
	name string
}

func NewGifter(id, name string) Gifter {
	return &gifter{
		id:   id,
		name: name,
	}
}

func (g *gifter) Name() string {
	return g.name
}

func (g *gifter) ID() string {
	return g.id
}

type group struct {
	id      string
	name    string
	budget  float64
	gifters []Gifter
}

func NewGroup(id, name string, budget float64, gifters []Gifter) Group {
	return &group{
		id:      id,
		name:    name,
		budget:  budget,
		gifters: gifters,
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
	for _, member := range grp.gifters {
		if member.ID() == g.ID() {
			return ErrGifterAlreadyInGroup
		}
	}
	grp.gifters = append(grp.gifters, g)
	return nil
}
