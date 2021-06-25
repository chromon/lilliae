package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 比较指令可以分为两类：
// 一类将比较结果推入操作数栈顶，一类根据比较结果跳转。
// 比较指令是编译器实现 if-else、for、while 等语句的基石

// fcmpg 和 fcmpl 指令用于比较 float 变量
// 由于浮点数计算有可能产生 NaN（Not a Number）值，所以比较两个浮点数时，
// 除了大于、等于、小于之外，还有第 4 种结果：无法比较。
// fcmpg 和 fcmpl 指令的区别就在于对第 4 种结果的定义

type FCMPG struct {
	base.NoOperandsInstruction
}

func (fc *FCMPG) Execute(frame *runtimedataarea.Frame) {
	// 当两个 float 变量中至少有一个是 NaN 时
	// 用 fcmpg 指令比较的结果是 1
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (fc *FCMPL) Execute(frame *runtimedataarea.Frame) {
	// 当两个 float 变量中至少有一个是 NaN 时
	// 用 fcmpl 指令比较的结果是 -1
	_fcmp(frame, false)
}

func _fcmp(frame *runtimedataarea.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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