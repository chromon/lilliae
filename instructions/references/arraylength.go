package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// arraylength 指令用于获取数组长度
type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (l *ARRAY_LENGTH) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	// arraylength 指令只需要一个操作数，即从操作数栈顶弹出的数组引用
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	// 将数组长度推入操作数栈顶
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
