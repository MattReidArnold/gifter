package entities

// Gifter is a participant in a gifting circle
type Gifter struct {
	Name          string
	DoNotGiftFrom []string
	ID            string
	GiftTo        string
}

// Circle is a group of gifters who will exchange gifts with each other
type Circle struct {
	ID      string
	Gifters []Gifter
}
