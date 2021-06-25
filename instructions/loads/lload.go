package loads

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 从局部变量表获取 long 变量，然后推入操作数栈顶
type LLOAD struct {
	base.Index8Instruction
}

func _lload(frame *runtimedataarea.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

func (ll *LLOAD) Execute(frame *runtimedataarea.Frame) {
	_lload(frame, ll.Index)
}

// 从局部变量表获取 long 变量 0，然后推入操作数栈顶
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (ll *LLOAD_0) Execute(frame *runtimedataarea.Frame) {
	_lload(frame, 0)
}

// 从局部变量表获取 long 变量 1，然后推入操作数栈顶
type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (ll *LLOAD_1) Execute(frame *runtimedataarea.Frame) {
	_lload(frame, 1)
}

// 从局部变量表获取 long 变量 2，然后推入操作数栈顶
type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (ll *LLOAD_2) Execute(frame *runtimedataarea.Frame) {
	_lload(frame, 2)
}

// 从局部变量表获取 long 变量 3，然后推入操作数栈顶
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (ll *LLOAD_3) Execute(frame *runtimedataarea.Frame) {
	_lload(frame, 3)
}