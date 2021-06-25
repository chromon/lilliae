package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 比较指令可以分为两类：
// 一类将比较结果推入操作数栈顶，一类根据比较结果跳转。
// 比较指令是编译器实现 if-else、for、while 等语句的基石

// dcmpg 和 dcmpl 指令用于比较 double 变量
type DCMPG struct {
	base.NoOperandsInstruction
}

func (dc *DCMPG) Execute(frame *runtimedataarea.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (dc *DCMPL) Execute(frame *runtimedataarea.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *runtimedataarea.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}