package presenters

type AddGifterRequest struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
}

type AddGifterResponse struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
}
