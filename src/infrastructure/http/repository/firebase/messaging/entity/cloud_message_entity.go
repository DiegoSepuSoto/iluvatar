package entity

type CloudMessageEntity struct {
	Notification *Notification `json:"notification"`
	SendTo       string       `json:"to"`
}

type Notification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
