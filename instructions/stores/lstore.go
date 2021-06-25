package stores

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 存储指令把变量从操作数栈顶弹出，然后存入局部变量表

// 把 long 型变量从操作数栈顶弹出，然后存入局部变量表
type LSTORE struct {
	base.Index8Instruction
}

func _lstore(frame *runtimedataarea.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (ls *LSTORE) Execute(frame *runtimedataarea.Frame) {
	_lstore(frame, ls.Index)
}

// 把 long 型变量 0 从操作数栈顶弹出，然后存入局部变量表
type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (ls *LSTORE_0) Execute(frame *runtimedataarea.Frame) {
	_lstore(frame, 0)
}

// 把 long 型变量 1 从操作数栈顶弹出，然后存入局部变量表
type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (ls *LSTORE_1) Execute(frame *runtimedataarea.Frame) {
	_lstore(frame, 1)
}

// 把 long 型变量 2 从操作数栈顶弹出，然后存入局部变量表
type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (ls *LSTORE_2) Execute(frame *runtimedataarea.Frame) {
	_lstore(frame, 2)
}

// 把 long 型变量 3 从操作数栈顶弹出，然后存入局部变量表
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (ls *LSTORE_3) Execute(frame *runtimedataarea.Frame) {
	_lstore(frame, 3)
}