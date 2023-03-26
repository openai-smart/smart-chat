package chat

import sc "github.com/openai-smart/smart-chat"

const MessageFromSoftware = "software"

// Filter 拦截器
type Filter interface {
	DoFilter(*sc.Session) error
}
