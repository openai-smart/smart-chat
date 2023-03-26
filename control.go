package smartchat

type Instruction string

type Command interface {
	Help() string
	Execute(...string) error
}

// Control 内置管理指令
type Control interface {
	Help() string
	AddCommand(Instruction, Command) error
}
