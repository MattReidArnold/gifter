package presenters

type AddGifterRequest struct {
	CircleID string `json:"circleId"`
	Name     string `json:"name"`
}

type AddGifterResponse struct {
	CircleID string `json:"circleId"`
	Name     string `json:"name"`
}
