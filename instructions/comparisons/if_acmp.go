package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// if_acmpeq 和 if_acmpne 指令把栈顶的两个引用弹出，根据引用是否相同进行跳转

// 栈顶的两个引用相同
type IF_ACMPEQ struct{
	base.BranchInstruction
}

func (eq *IF_ACMPEQ) Execute(frame *runtimedataarea.Frame) {
	if _acmp(frame) {
		base.Branch(frame, eq.Offset)
	}
}

// 栈顶的两个引用不同
type IF_ACMPNE struct{
	base.BranchInstruction
}

func (ne *IF_ACMPNE) Execute(frame *runtimedataarea.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, ne.Offset)
	}
}

func _acmp(frame *runtimedataarea.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
