package math

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 取反指令

// 取反指令 double
type DNEG struct {
	base.NoOperandsInstruction
}

func (dn *DNEG) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// 取反指令 float
type FNEG struct {
	base.NoOperandsInstruction
}

func (fn *FNEG) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// 取反指令 int
type INEG struct {
	base.NoOperandsInstruction
}

func (in *INEG) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// 取反指令 long
type LNEG struct {
	base.NoOperandsInstruction
}

func (ln *LNEG) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}