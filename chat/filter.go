package chat

import (
	"fmt"
	sc "github.com/openai-smart/smart-chat"
	"github.com/pkg/errors"
	"time"
)

// Filter 拦截器，如果err!=nil表示需要拦截此请求
type Filter interface {
	AccessFilter(user User) error                                                                      // 用户拦截器
	IncomingMessageFilter(msg string, msgUid string, sendTime time.Time, msgType sc.MessageType) error // 消息拦截器
}

type DefaultFilter struct {
	Filter

	// TODO 接入数据源
}

func (filter *DefaultFilter) IncomingMessageFilter(msg string, msgUid string, sendTime time.Time, msgType sc.MessageType) error {
	if msgType != sc.MessageTypeText { // 拦截消息不是MessageTypeText类型的
		return errors.New(string(fmt.Sprintf("message type [%s] unsupported", msgType)))
	}

	t := time.Now().Sub(sendTime)
	if t > 10*time.Second { // 拦截丢弃超时消息
		return errors.New(fmt.Sprintf("overtime %d, %s ", t, msg))
	}

	// TODO uid 拦截重复消息
	return nil
}

func (filter *DefaultFilter) AccessFilter(user User) error {
	// TODO
	return nil
}
