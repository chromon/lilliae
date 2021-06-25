package stack

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// swap 指令交换栈顶的两个变量
/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
type SWAP struct {
	base.NoOperandsInstruction
}

func (s *SWAP) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}