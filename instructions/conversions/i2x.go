package conversions

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 强制转换指令：将 int 类型变量强制转换为其他类型

// 将 int 类型强制转换为 byte 类型
type I2B struct {
	base.NoOperandsInstruction
}

func (ib *I2B) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// 将 int 类型强制转换为 char 类型
type I2C struct {
	base.NoOperandsInstruction
}

func (ic *I2C) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// 将 int 类型强制转换为 short 类型
type I2S struct {
	base.NoOperandsInstruction
}

func (is *I2S) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// 将 int 类型强制转换为 long 类型
type I2L struct {
	base.NoOperandsInstruction
}

func (il *I2L) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// 将 int 类型强制转换为 float 类型
type I2F struct {
	base.NoOperandsInstruction
}

func (i2f *I2F) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// 将 int 类型强制转换为 double 类型
type I2D struct {
	base.NoOperandsInstruction
}

func (id *I2D) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}