package loads

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 加载指令从局部变量表获取变量，然后推入操作数栈顶

// 操作 int 类型变量
type ILOAD struct {
	base.Index8Instruction
}

func _iload(frame *runtimedataarea.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (il *ILOAD) Execute(frame *runtimedataarea.Frame) {
	_iload(frame, uint(il.Index))
}

// 从局部变量表中获取变量 int 型变量 0，压入操作数栈顶
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (il *ILOAD_0) Execute(frame *runtimedataarea.Frame) {
	_iload(frame, 0)
}

// 从局部变量表中获取变量 int 型变量 1，压入操作数栈顶
type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (il *ILOAD_1) Execute(frame *runtimedataarea.Frame) {
	_iload(frame, 1)
}

// 从局部变量表中获取变量 int 型变量 2，压入操作数栈顶
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (il *ILOAD_2) Execute(frame *runtimedataarea.Frame) {
	_iload(frame, 2)
}

// 从局部变量表中获取变量 int 型变量 3，压入操作数栈顶
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (il *ILOAD_3) Execute(frame *runtimedataarea.Frame) {
	_iload(frame, 3)
}