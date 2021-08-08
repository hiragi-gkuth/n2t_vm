package vm

// Command Types
type CommandType uint8

const (
	C_ARTITHMETIC CommandType = iota
	C_PUSH
	C_POP
	C_LABEL
	C_GOTO
	C_IF
	C_FUNCTION
	C_RETURN
	C_CALL
)

// Offsets of RAM
const (
	O_REGISTER = 0x00000000
	O_STATIC   = 0x00000010
	O_STACK    = 0x00000100
	O_HEAP     = 0x00001000
	O_IOMAP    = 0x05000000
)

// Instruction Types
type Instruction string

const (
	I_ADD      Instruction = "add"
	I_SUB      Instruction = "sub"
	I_NEG      Instruction = "neg"
	I_EQ       Instruction = "eq"
	I_GT       Instruction = "gt"
	I_LT       Instruction = "lt"
	I_AND      Instruction = "and"
	I_OR       Instruction = "or"
	I_NOT      Instruction = "not"
	I_PUSH     Instruction = "push"
	I_POP      Instruction = "pop"
	I_LABEL    Instruction = "label"
	I_IF_GOTO  Instruction = "if-goto"
	I_FUNCTION Instruction = "function"
	I_RETURN   Instruction = "return"
	I_CALL     Instruction = "call"
)
