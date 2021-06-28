package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (is *INVOKE_SPECIAL) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PopRef()
}