package smartchat

import (
	"time"
)

type Question any
type Answer any
type SessionSource int
type SessionID string
type SessionStatus int

const (
	// SessionStatusNone 会话不存在
	SessionStatusNone SessionStatus = iota
	// SessionStatusProcessing 会话处理中
	SessionStatusProcessing
	// SessionStatusCompletion 会话处理完成
	SessionStatusCompletion
	// SessionStatusErr 会话失败
	SessionStatusErr
)

// SessionError 对话失败错误信息
type SessionError struct {
	MsgID   string // 消息ID
	ErrCode int    // 错误代码
	Message string // 错误关键信息
	Error   error  // 错误堆栈信息
}

// Session 会话
type Session struct {
	// 会话ID
	ID SessionID `json:"id"`
	// 会话来源
	Source SessionSource `json:"source"`
	// 用户
	User *User `json:"user"`
	// smart id
	SmartID string `json:"smart_id"`
	// 问题
	Question Question `json:"question"`
	// 答案
	Answer Answer `json:"answer"`
	// 对话开始时间
	Start time.Time `json:"start"`
	// 对话结束时间
	End time.Time `json:"end"`
	// 对话是否成功
	Complete bool `json:"complete"`
}

// MessageType 消息类型
type MessageType string

// MessageTypeText 文本消息
const MessageTypeText MessageType = "text"

// MessageTypeVoice 语音消息
const MessageTypeVoice MessageType = "voice"

// MessageTypeImage 图片消息
const MessageTypeImage MessageType = "image"

// MessageTypeVideo 视频消息
const MessageTypeVideo MessageType = "video"

// MessageTypeUnknown 未知消息类型
const MessageTypeUnknown MessageType = "text"
