package stack

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 将栈顶变量弹出
type POP struct {
	base.NoOperandsInstruction
}

// POP 指令只能用于弹出 int、float 等占用一个操作数栈位置的变量
func (p *POP) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 将栈顶变量弹出
type POP2 struct {
	base.NoOperandsInstruction
}

// POP2 指令可以用于弹出 double 和 long 类型占用两个操作数栈位置的变量
func (p *POP2) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}