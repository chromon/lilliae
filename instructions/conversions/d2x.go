package conversions

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 强制转换指令：将 double 类型变量强制转换为其他类型

// 将 double 类型强制转换为 int 类型
type D2I struct {
	base.NoOperandsInstruction
}

func (di *D2I) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// 将 double 类型强制转换为 float 类型
type D2F struct {
	base.NoOperandsInstruction
}

func (df *D2F) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// 将 double 类型强制转换为 long 类型
type D2L struct {
	base.NoOperandsInstruction
}

func (dl *D2L) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}