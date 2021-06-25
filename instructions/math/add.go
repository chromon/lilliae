package math

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 加法指令

// double 加法
type DADD struct {
	base.NoOperandsInstruction
}

func (da *DADD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// float 加法指令
type FADD struct {
	base.NoOperandsInstruction
}

func (fa *FADD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// int 加法指令
type IADD struct {
	base.NoOperandsInstruction
}

func (ia *IADD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// long 加法指令
type LADD struct {
	base.NoOperandsInstruction
}

func (la *LADD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}