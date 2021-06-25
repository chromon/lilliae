package loads

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 从局部变量表获取 double 变量，然后推入操作数栈顶
type DLOAD struct {
	base.Index8Instruction
}

func _dload(frame *runtimedataarea.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (dl *DLOAD) Execute(frame *runtimedataarea.Frame) {
	_dload(frame, dl.Index)
}

// 从局部变量表获取 double 变量 0，然后推入操作数栈顶
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (dl *DLOAD_0) Execute(frame *runtimedataarea.Frame) {
	_dload(frame, 0)
}

// 从局部变量表获取 double 变量 1，然后推入操作数栈顶
type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (dl *DLOAD_1) Execute(frame *runtimedataarea.Frame) {
	_dload(frame, 1)
}

// 从局部变量表获取 double 变量 2，然后推入操作数栈顶
type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (dl *DLOAD_2) Execute(frame *runtimedataarea.Frame) {
	_dload(frame, 2)
}

// 从局部变量表获取 double 变量 3，然后推入操作数栈顶
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (dl *DLOAD_3) Execute(frame *runtimedataarea.Frame) {
	_dload(frame, 3)
}