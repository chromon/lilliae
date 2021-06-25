package conversions

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 强制转换指令：将 long 类型变量强制转换为其他类型

// 将 long 类型强制转换为 double 类型
type L2D struct {
	base.NoOperandsInstruction
}

func (ld *L2D) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// 将 long 类型强制转换为 float 类型
type L2F struct {
	base.NoOperandsInstruction
}

func (lf *L2F) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// 将 long 类型强制转换为 int 类型
type L2I struct {
	base.NoOperandsInstruction
}

func (li *L2I) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}