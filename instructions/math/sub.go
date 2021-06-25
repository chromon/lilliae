package math

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 减法指令

// 减法指令 double
type DSUB struct {
	base.NoOperandsInstruction
}

func (ds *DSUB) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// 减法指令 float
type FSUB struct {
	base.NoOperandsInstruction
}

func (fs *FSUB) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// 减法指令 int
type ISUB struct {
	base.NoOperandsInstruction
}

func (is *ISUB) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// 减法指令 long
type LSUB struct {
	base.NoOperandsInstruction
}

func (ls *LSUB) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}