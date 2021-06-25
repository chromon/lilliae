package stores

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 把 int 型变量从操作数栈顶弹出，然后存入局部变量表
type ISTORE struct {
	base.Index8Instruction
}

func _istore(frame *runtimedataarea.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (is *ISTORE) Execute(frame *runtimedataarea.Frame) {
	_istore(frame, uint(is.Index))
}

// 把 int 型变量 0 从操作数栈顶弹出，然后存入局部变量表
type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (is *ISTORE_0) Execute(frame *runtimedataarea.Frame) {
	_istore(frame, 0)
}

// 把 int 型变量 1 从操作数栈顶弹出，然后存入局部变量表
type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (is *ISTORE_1) Execute(frame *runtimedataarea.Frame) {
	_istore(frame, 1)
}

// 把 int 型变量 2 从操作数栈顶弹出，然后存入局部变量表
type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (is *ISTORE_2) Execute(frame *runtimedataarea.Frame) {
	_istore(frame, 2)
}

// 把 int 型变量 3 从操作数栈顶弹出，然后存入局部变量表
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (is *ISTORE_3) Execute(frame *runtimedataarea.Frame) {
	_istore(frame, 3)
}