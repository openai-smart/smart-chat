package smart

import (
	"github.com/google/uuid"
	sc "github.com/openai-smart/smart-chat"
)

// CompletionHandler 回复消息处理函数
type CompletionHandler func(msg *sc.Message) error

// Smart 人工智能
type Smart interface {

	//AddCompletionHandler 新增回复信息处理函数，注意这是在原来基础上新增
	AddCompletionHandler(CompletionHandler)

	//SetCompletionHandler 清空回复信息处理函数，添加此次设置的函数
	SetCompletionHandler(CompletionHandler)

	//Question 提问
	Question(*sc.Message) error

	// Balance 余额
	Balance() int

	// UUID 唯一ID
	UUID() uuid.UUID
}
