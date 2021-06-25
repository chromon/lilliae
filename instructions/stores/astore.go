package stores

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 把 引用 型变量从操作数栈顶弹出，然后存入局部变量表
type ASTORE struct {
	base.Index8Instruction
}

func _astore(frame *runtimedataarea.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

func (as *ASTORE) Execute(frame *runtimedataarea.Frame) {
	_astore(frame, uint(as.Index))
}

// 把 引用 型变量 0 从操作数栈顶弹出，然后存入局部变量表
type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (as *ASTORE_0) Execute(frame *runtimedataarea.Frame) {
	_astore(frame, 0)
}

// 把 引用 型变量 1 从操作数栈顶弹出，然后存入局部变量表
type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (as *ASTORE_1) Execute(frame *runtimedataarea.Frame) {
	_astore(frame, 1)
}

// 把 引用 型变量 2 从操作数栈顶弹出，然后存入局部变量表
type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (as *ASTORE_2) Execute(frame *runtimedataarea.Frame) {
	_astore(frame, 2)
}

// 把 引用 型变量 3 从操作数栈顶弹出，然后存入局部变量表
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (as *ASTORE_3) Execute(frame *runtimedataarea.Frame) {
	_astore(frame, 3)
}