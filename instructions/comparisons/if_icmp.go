package comparisons

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// f_icmp<cond> 指令把栈顶的两个 int 变量弹出，然后进行比较，满足条件则跳转

// 当弹出的栈顶两个变量相等时跳转
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (eq *IF_ICMPEQ) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, eq.Offset)
	}
}

// 当弹出的栈顶两个变量不相等时跳转
type IF_ICMPNE struct {
	base.BranchInstruction
}

func (ne *IF_ICMPNE) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, ne.Offset)
	}
}

// 当弹出的栈顶变量 v2 比下一个栈顶元素 v1 大时跳转
type IF_ICMPLT struct {
	base.BranchInstruction
}

func (lt *IF_ICMPLT) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, lt.Offset)
	}
}

// 当弹出的栈顶变量 v2 大于等于下一个栈顶元素 v1 时跳转
type IF_ICMPLE struct {
	base.BranchInstruction
}

func (le *IF_ICMPLE) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, le.Offset)
	}
}

// 当弹出的栈顶变量 v2 小于下一个栈顶元素 v1 时跳转
type IF_ICMPGT struct {
	base.BranchInstruction
}

func (gt *IF_ICMPGT) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, gt.Offset)
	}
}

// 当弹出的栈顶变量 v2 小于等于下一个栈顶元素 v1 时跳转
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (ge *IF_ICMPGE) Execute(frame *runtimedataarea.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, ge.Offset)
	}
}

func _icmpPop(frame *runtimedataarea.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}