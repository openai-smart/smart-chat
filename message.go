package smartchat

import (
	"github.com/google/uuid"
)

type Original struct {
	Question any
	Answer   any
}

// Message 消息结构
type Message struct {
	UserID    string
	SmartUUID uuid.UUID
	Original  Original // 原始消息
}

// MessageType 消息类型
type MessageType string

// MessageTypeUnknown 未知消息类型
const MessageTypeUnknown MessageType = "text"

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeLocation 位置消息
const MessageTypeLocation MessageType = "location"

// MessageTypeLink 链接消息
const MessageTypeLink MessageType = "link"

// MessageTypeEvent 事件消息
const MessageTypeEvent MessageType = "event"
