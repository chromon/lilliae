package extended

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 根据引用是否是 null 进行跳转

// 引用是 null
type IFNULL struct {
	base.BranchInstruction
}

func (in *IFNULL) Execute(frame *runtimedataarea.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, in.Offset)
	}
}

// 引用不是 null
type IFNONNULL struct {
	base.BranchInstruction
}

func (inn *IFNONNULL) Execute(frame *runtimedataarea.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, inn.Offset)
	}
}