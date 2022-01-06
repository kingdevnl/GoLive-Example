package message

import "time"

type Message struct {
	User    string
	Message string
	Date    string
}

var Messages = []*Message{
	NewMessage("James", "Hello World"),
	NewMessage("John", "Hello James"),
	NewMessage("James", "Hello John"),
}

func NewMessage(user string, message string) *Message {
	return &Message{
		User:    user,
		Message: message,
		Date:    time.Now().Format(time.RFC3339),
	}
}
