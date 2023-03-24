package smartchat

type Cache interface {
	Session(msg Message)
}
