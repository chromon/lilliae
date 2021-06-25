package constants

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 常量指令把常量推入操作数栈顶
// 常量可以来自三个地方：隐含在操作码里、操作数和运行时常量池

// nop 指令是最简单的一条指令，因为它什么都不做
type NOP struct {
	base.NoOperandsInstruction
}

// 执行指令逻辑什么都不做
func (nop *NOP) Execute(frame *runtimedataarea.Frame) {}