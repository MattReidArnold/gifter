package entities

// Gifter is a participant in a gifting circle
type Gifter struct {
	Name          string
	DoNotGiftFrom []string
	ID            string
	GiftTo        string
}
