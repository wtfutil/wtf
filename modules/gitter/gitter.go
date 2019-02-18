package gitter

import "time"

type Rooms struct {
	Results []Room `json:"results"`
}

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URI  string `json:"uri"`
}

type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type Message struct {
	ID     string    `json:"id"`
	Text   string    `json:"text"`
	HTML   string    `json:"html"`
	Sent   time.Time `json:"sent"`
	From   User      `json:"fromUser"`
	Unread bool      `json:"unread"`
}
