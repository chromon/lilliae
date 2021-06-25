package base

import "lilliae/runtimedataarea"

// 指令接口
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *runtimedataarea.Frame)
}

// 没有操作数的指令，没有定义任何字段
type NoOperandsInstruction struct {}

// 提取操作数方法也没有内容
func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	// 跳转偏移量
	Offset int
}

// 从字节码中提取操作数
func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码中读取 uint16 整数
	bi.Offset = int(reader.ReadInt16())
}

// 存储和加载类指令需要根据索引存取局部变量表，索引由单节操作数给出。
// 这类指令抽象为 Index8Instruction 结构体，用 Index 字段表示局部变量表索引
type Index8Instruction struct {
	// 局部变量表索引
	Index uint
}

// 从字节码中读取操作数
func (ii *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	ii.Index = uint(reader.ReadUint8())
}

// 一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。
// 把这类指令抽象成 Index16Instruction 结构体，
// 用 Index 字段表示常量池索引
type Index16Instruction struct {
	// 常量池索引
	Index uint
}

// 从字节码中读取操作数
func (ii *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	ii.Index = uint(reader.ReadUint16())
}