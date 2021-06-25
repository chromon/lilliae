package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 比较指令可以分为两类：
// 一类将比较结果推入操作数栈顶，一类根据比较结果跳转。
// 比较指令是编译器实现 if-else、for、while 等语句的基石

// lcmp 指令用于比较 long 变量
type LCMP struct {
	base.NoOperandsInstruction
}

// 弹出，比较，将最终结果推入栈顶
func (lc *LCMP) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}