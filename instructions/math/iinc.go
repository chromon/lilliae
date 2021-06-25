package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// iinc 指令给局部变量表中的 int 变量增加常量值
// 局部变量表索引和常量值都有指令的操作数提供
type IINC struct {
	Index uint
	Const int32
}

// 从字节码中读取操作数
func (ii *IINC) FetchOperands(reader *base.BytecodeReader) {
	ii.Index = uint(reader.ReadUint8())
	ii.Const = int32(reader.ReadInt8())
}

// 从局部变量表中读取变量，加上常量值，再将结果写回到局部变量表
func (ii *IINC) Execute(frame *runtimedataarea.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(ii.Index)
	val += ii.Const
	localVars.SetInt(ii.Index, val)
}