package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// if<cond> 指令把操作数栈顶的 int 变量弹出，然后跟 0 进行比较，满足条件则跳转

// 当弹出的栈顶变量等于 0 时跳转
type IFEQ struct {
	base.BranchInstruction
}

func (eq *IFEQ) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, eq.Offset)
	}
}

// 当弹出的栈顶变量不等于 0 时跳转
type IFNE struct {
	base.BranchInstruction
}

func (ne *IFNE) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, ne.Offset)
	}
}

// 当弹出的栈顶变量小于 0 时跳转
type IFLT struct {
	base.BranchInstruction
}

func (lt *IFLT) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, lt.Offset)
	}
}

// 当弹出的栈顶变量小于等于 0 时跳转
type IFLE struct {
	base.BranchInstruction
}

func (le *IFLE) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, le.Offset)
	}
}

// 当弹出的栈顶变量大于 0 时跳转
type IFGT struct {
	base.BranchInstruction
}

func (gt *IFGT) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, gt.Offset)
	}
}

// 当弹出的栈顶变量大于等于 0 时跳转
type IFGE struct {
	base.BranchInstruction
}

func (ge *IFGE) Execute(frame *runtimedataarea.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, ge.Offset)
	}
}