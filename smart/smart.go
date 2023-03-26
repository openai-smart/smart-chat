package smart

import (
	sc "github.com/openai-smart/smart-chat"
)

// Configure AI配置信息
type Configure any

// Smart 人工智能/AI
type Smart interface {
	// Platform 平台
	Platform() string

	// Ask 提问
	Ask(sc.Question) (sc.Answer, error)

	// Balance 余额
	Balance() (float32, error)
}
