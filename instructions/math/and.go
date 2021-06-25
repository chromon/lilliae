package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 按位与
// 布尔运算只能操作 int 和 long 变量

// int 型按位与运算
type IAND struct {
	base.NoOperandsInstruction
}

func (ia *IAND) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// long 型按位与运算
type LAND struct {
	base.NoOperandsInstruction
}

func (la *LAND) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}