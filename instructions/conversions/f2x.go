package conversions

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 强制转换指令：将 float 类型变量强制转换为其他类型

// 将 float 类型强制转换为 int 类型
type F2I struct {
	base.NoOperandsInstruction
}

func (fi *F2I) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// 将 float 类型强制转换为 double 类型
type F2D struct {
	base.NoOperandsInstruction
}

func (fd *F2D) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// 将 float 类型强制转换为 long 类型
type F2L struct {
	base.NoOperandsInstruction
}

func (fl *F2L) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}