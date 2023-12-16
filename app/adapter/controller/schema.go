package controller

type Header struct {
	RoomID       string `json:"room_id"`
	WantClientID string `json:"want_client_id"`
}
type Body struct {
	Content string `json:"content"`
}
type Protocol struct {
	Type   string `json:"type"`
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}
