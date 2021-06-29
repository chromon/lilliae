package control

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
)

// 返回指令

// 没有返回值
type RETURN struct {
	base.NoOperandsInstruction
}

func (r *RETURN) Execute(frame *Frame) {
	frame.Thread().PopFrame()
}

// 返回引用
type ARETURN struct {
	base.NoOperandsInstruction
}

func (a *ARETURN) Execute(frame *Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

// 返回 double
type DRETURN struct {
	base.NoOperandsInstruction
}

func (d *DRETURN) Execute(frame *Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

// 返回 float
type FRETURN struct {
	base.NoOperandsInstruction
}

func (f *FRETURN) Execute(frame *Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

// 返回 int
type IRETURN struct {
	base.NoOperandsInstruction
}

func (i *IRETURN) Execute(frame *Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

// 返回 long
type LRETURN struct {
	base.NoOperandsInstruction
}

func (l *LRETURN) Execute(frame *Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}