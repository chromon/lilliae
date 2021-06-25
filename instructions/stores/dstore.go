package stores

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 把 double 型变量从操作数栈顶弹出，然后存入局部变量表
type DSTORE struct {
	base.Index8Instruction
}

func _dstore(frame *runtimedataarea.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (ds *DSTORE) Execute(frame *runtimedataarea.Frame) {
	_dstore(frame, uint(ds.Index))
}

// 把 double 型变量 0 从操作数栈顶弹出，然后存入局部变量表
type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (ds *DSTORE_0) Execute(frame *runtimedataarea.Frame) {
	_dstore(frame, 0)
}

// 把 double 型变量 1 从操作数栈顶弹出，然后存入局部变量表
type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (ds *DSTORE_1) Execute(frame *runtimedataarea.Frame) {
	_dstore(frame, 1)
}

// 把 double 型变量 2 从操作数栈顶弹出，然后存入局部变量表
type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (ds *DSTORE_2) Execute(frame *runtimedataarea.Frame) {
	_dstore(frame, 2)
}

// 把 double 型变量 3 从操作数栈顶弹出，然后存入局部变量表
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (ds *DSTORE_3) Execute(frame *runtimedataarea.Frame) {
	_dstore(frame, 3)
}