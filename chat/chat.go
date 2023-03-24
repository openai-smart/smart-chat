package chat

import sc "github.com/openai-smart/smart-chat"

type Chat interface {
	Reply(msg *sc.Message) error
	Accept(string) error
	SetFilter(Filter)
}

type User struct {
	ID    string
	Name  string
	Roles []int
}
