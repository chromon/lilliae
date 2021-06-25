package loads

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 从局部变量表获取 float 变量，然后推入操作数栈顶
type FLOAD struct {
	base.Index8Instruction
}

func _fload(frame *runtimedataarea.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (fl *FLOAD) Execute(frame *runtimedataarea.Frame) {
	_fload(frame, fl.Index)
}

// 从局部变量表获取 float 变量 0，然后推入操作数栈顶
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (fl *FLOAD_0) Execute(frame *runtimedataarea.Frame) {
	_fload(frame, 0)
}

// 从局部变量表获取 float 变量 1，然后推入操作数栈顶
type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (fl *FLOAD_1) Execute(frame *runtimedataarea.Frame) {
	_fload(frame, 1)
}

// 从局部变量表获取 float 变量 2，然后推入操作数栈顶
type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (fl *FLOAD_2) Execute(frame *runtimedataarea.Frame) {
	_fload(frame, 2)
}

// 从局部变量表获取 float 变量 3，然后推入操作数栈顶
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (fl *FLOAD_3) Execute(frame *runtimedataarea.Frame) {
	_fload(frame, 3)
}