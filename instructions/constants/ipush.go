package constants

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 从操作数中获取一个 byte 型正数，扩展为 int 型，后压入栈顶
type BIPUSH struct {
	val int8
}

// 从字节码中提取操作数
func (bip *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	bip.val = reader.ReadInt8()
}

// 执行指令逻辑
func (bip *BIPUSH) Execute(frame *runtimedataarea.Frame) {
	i := int32(bip.val)
	frame.OperandStack().PushInt(i)
}

// 从操作数中获取一个 short 型整数，扩展成 int 型，然后推入栈顶
type SIPUSH struct {
	val int16
}

// 从字节码中提取操作数
func (sip *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	sip.val = reader.ReadInt16()
}

// 执行指令逻辑
func (sip *SIPUSH) Execute(frame *runtimedataarea.Frame) {
	i := int32(sip.val)
	frame.OperandStack().PushInt(i)
}