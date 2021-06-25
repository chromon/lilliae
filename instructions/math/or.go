package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 按位或
// 布尔运算只能操作 int 和 long 变量

// int 型按位或运算
type IOR struct{
	base.NoOperandsInstruction
}

func (io *IOR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// long 型按位或运算
type LOR struct{
	base.NoOperandsInstruction
}

func (lo *LOR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}