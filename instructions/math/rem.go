package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"math"
)

// 求余指令

// double 类型求余
type DREM struct {
	base.NoOperandsInstruction
}

func (dr *DREM) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	// 浮点数类型因为有 Infinity（无穷大） 值，所以即使是除零，
	// 也不会导致 ArithmeticException 异常抛出
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

// float 类型求余
type FREM struct {
	base.NoOperandsInstruction
}

func (fr *FREM) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	// 求余
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

// int 类型求余
type IREM struct {
	base.NoOperandsInstruction
}

func (ir *IREM) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// long 类型求余
type LREM struct {
	base.NoOperandsInstruction
}

func (lr *LREM) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}