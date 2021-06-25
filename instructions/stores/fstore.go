package stores

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 把 float 型变量从操作数栈顶弹出，然后存入局部变量表
type FSTORE struct {
	base.Index8Instruction
}

func _fstore(frame *runtimedataarea.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (fs *FSTORE) Execute(frame *runtimedataarea.Frame) {
	_fstore(frame, uint(fs.Index))
}

// 把 float 型变量 0 从操作数栈顶弹出，然后存入局部变量表
type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (fs *FSTORE_0) Execute(frame *runtimedataarea.Frame) {
	_fstore(frame, 0)
}

// 把 float 型变量 1 从操作数栈顶弹出，然后存入局部变量表
type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (fs *FSTORE_1) Execute(frame *runtimedataarea.Frame) {
	_fstore(frame, 1)
}

// 把 float 型变量 2 从操作数栈顶弹出，然后存入局部变量表
type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (fs *FSTORE_2) Execute(frame *runtimedataarea.Frame) {
	_fstore(frame, 2)
}

// 把 float 型变量 3 从操作数栈顶弹出，然后存入局部变量表
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (fs *FSTORE_3) Execute(frame *runtimedataarea.Frame) {
	_fstore(frame, 3)
}