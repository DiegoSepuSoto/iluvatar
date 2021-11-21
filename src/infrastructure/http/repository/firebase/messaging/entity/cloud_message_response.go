package entity

type CloudMessageResponse struct {
	MulticastID  int64    `json:"multicast_id"`
	Success      int64    `json:"success"`
	Failure      int64    `json:"failure"`
	CanonicalIDS int64    `json:"canonical_ids"`
	Results      []*Result `json:"results"`
}

type Result struct {
	MessageID string `json:"message_id"`
}
