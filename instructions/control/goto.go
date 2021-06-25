package control

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// goto 指令进行无条件跳转
type GOTO struct {
	base.BranchInstruction
}

func (gt *GOTO) Execute(frame *runtimedataarea.Frame) {
	base.Branch(frame, gt.Offset)
}
