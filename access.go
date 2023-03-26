package smartchat

// Role 角色
type Role int

type UserUID string // 用户唯一ID
type UserStatus int // 用户状态

const (
	// UserStatusActive 激活状态
	UserStatusActive UserStatus = iota
	// UserStatusDown 由于某些原因导致的账户暂时不可用
	UserStatusDown
	// UserStatusBlock 账户已被锁定
	UserStatusBlock
)

type User struct {
	UID    UserUID
	Name   string
	Status UserStatus
	Roles  []Role
}
