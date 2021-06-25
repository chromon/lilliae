package math

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 乘法指令

// 乘法指令 double
type DMUL struct {
	base.NoOperandsInstruction
}

func (dm *DMUL) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

// 乘法指令 float
type FMUL struct {
	base.NoOperandsInstruction
}

func (fm *FMUL) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

// 乘法指令 int
type IMUL struct {
	base.NoOperandsInstruction
}

func (im *IMUL) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

// 乘法指令 long
type LMUL struct {
	base.NoOperandsInstruction
}

func (lm *LMUL) Execute(frame *Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
