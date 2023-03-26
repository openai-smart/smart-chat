package smartchat

import "time"

type Configure map[string]any

type Cache interface {
	// Error 读取错误信息
	Error(SessionID) (SessionError, error)

	// ErrorStore 错误信息
	ErrorStore(SessionID, SessionError) error

	// Session 取出对应ID的会话
	Session(SessionID) (*Session, error)

	// SessionStore 保存会话
	SessionStore(*Session) error

	// SessionStatusRecord 记录会话状态
	SessionStatusRecord(SessionID, SessionStatus) error

	// SessionStatus 会话状态
	SessionStatus(SessionID) (SessionStatus, error)

	// SessionHistory 聊天记录
	SessionHistory(UserUID, time.Time, time.Time) ([]Session, error)

	// User 取得用户信息
	User(UserUID) (*User, error)

	// UserID2UID 获取 UserUID
	UserID2UID(string) (UserUID, error)

	// UserUIDBind 不同平台用户ID绑定 UserUID
	UserUIDBind(string, UserUID) error

	// UserStore 存储用户
	UserStore(*User) error

	// UserSmartsStore 为用户分配 smart
	UserSmartsStore(UserUID, ...string) error

	// UserSmartUIDs 返回此用户“拥有”权限的 smart
	UserSmartUIDs(UserUID) ([]string, error)

	// UserAnswer 配置需要使用的 smart
	UserAnswer(UserUID) ([]string, error)

	// UserAnswerStore 一个问题会询问所有配置的smart
	// ...string 聊天的smart id切片
	UserAnswerStore(UserUID, ...string) error

	// UserBalance TODO  用户余额
	UserBalance(UserUID) (float32, error)

	//Configure 配置信息
	Configure(string) (Configure, error)

	// ConfigureStore 存储配置
	ConfigureStore(string, Configure) error
}
