package loads

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 从局部变量表获取引用变量，然后推入操作数栈顶
type ALOAD struct {
	base.Index8Instruction
}

func _aload(frame *runtimedataarea.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

func (al *ALOAD) Execute(frame *runtimedataarea.Frame) {
	_aload(frame, al.Index)
}

// 从局部变量表获取引用变量 0，然后推入操作数栈顶
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (al *ALOAD_0) Execute(frame *runtimedataarea.Frame) {
	_aload(frame, 0)
}

// 从局部变量表获取引用变量 1，然后推入操作数栈顶
type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (al *ALOAD_1) Execute(frame *runtimedataarea.Frame) {
	_aload(frame, 1)
}

// 从局部变量表获取引用变量 2，然后推入操作数栈顶
type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (al *ALOAD_2) Execute(frame *runtimedataarea.Frame) {
	_aload(frame, 2)
}

// 从局部变量表获取引用变量，然后推入操作数栈顶
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (al *ALOAD_3) Execute(frame *runtimedataarea.Frame) {
	_aload(frame, 3)
}