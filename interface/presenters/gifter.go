package presenters

type AddGifterRequest struct {
	Name string `json:"name"`
}

type AddGifterResponse struct {
	Name string `json:"name"`
}
