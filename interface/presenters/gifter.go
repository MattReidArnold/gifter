package presenters

type AddGifterRequest struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
}

type AddGifterResponse struct {
	GifterID string `json:"gifterId"`
	GroupID  string `json:"groupId"`
	Name     string `json:"name"`
}
