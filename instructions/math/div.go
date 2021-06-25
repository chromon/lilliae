package math

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 除法指令

// double 除法指令
type DDIV struct {
	base.NoOperandsInstruction
}

func (d *DDIV) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// float 除法指令
type FDIV struct {
	base.NoOperandsInstruction
}

func (f *FDIV) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// int 除法指令
type IDIV struct {
	base.NoOperandsInstruction
}

func (i *IDIV) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// long 除法指令
type LDIV struct {
	base.NoOperandsInstruction
}

func (l *LDIV) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}