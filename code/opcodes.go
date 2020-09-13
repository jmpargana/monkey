package code

type Opcode byte

const (
	OpConstant Opcode = iota

	OpAdd // Integer Operators
	OpPop
	OpSub
	OpMul
	OpDiv

	OpTrue // Boolean Constants
	OpFalse

	OpEqual // Boolean Operators
	OpNotEqual
	OpGreaterThan

	OpMinus // Prefix Operators
	OpBang
)
