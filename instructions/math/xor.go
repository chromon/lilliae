package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 按位异或
// 布尔运算只能操作 int 和 long 变量

// int 型按位异或运算
type IXOR struct{
	base.NoOperandsInstruction
}

func (ixo *IXOR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// long 型按位异或运算
type LXOR struct{
	base.NoOperandsInstruction
}

func (lxo *LXOR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}