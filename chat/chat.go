package chat

import sc "github.com/openai-smart/smart-chat"

type EventConfigure any

// CompletionHandler 处理答复信息
type CompletionHandler func(*sc.Session) error

// Chat 聊天接口
type Chat interface {
	// Platform 平台
	Platform() string

	// AddCompletionHandler 增加处理答复信息方式，处理顺序先进先出
	AddCompletionHandler(CompletionHandler)

	// AddEventHandler 事件处理器
	AddEventHandler(EventConfigure) error

	// Accept 开始等待接收信息
	Accept(string) error

	// Filters 已配置的拦截器
	Filters() []Filter

	// AddFilter 新增拦截器
	AddFilter(Filter)
}
