package control

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// Java 语言中的 switch-case 语句有两种实现方式：
// 如果 case 值可以编码成一个索引表，则实现成 tableswitch 指令；
// 否则实现成 lookupswitch 指令

/*
tableswitch
<0-3 byte padding>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/

type TABLE_SWITCH struct {
	// 对应默认情况下执行跳转所需的字节码偏移量
	defaultOffset int32
	// case 的取值范围
	low int32
	// case 的取值范围
	high int32
	// 索引表，保存 high - low + 1 个 int 值
	// 对应各种 case 情况下，执行跳转所需要的字节码偏移量
	jumpOffsets []int32
}

func (ts *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// tableswitch 指令操作码的后面有 0 ~ 3 字节的 padding，
	// 以保证 defaultOffset 在字节码中的地址是 4 的倍数
	reader.SkipPadding()
	ts.defaultOffset = reader.ReadInt32()
	ts.low = reader.ReadInt32()
	ts.high = reader.ReadInt32()
	jumpOffsetsCount := ts.high - ts.low + 1
	ts.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (ts *TABLE_SWITCH) Execute(frame *runtimedataarea.Frame) {
	// 从操作数栈中弹出一个 int 变量
	index := frame.OperandStack().PopInt()
	var offset int
	// 是否是在 low 和 high 给定的范围内
	if index >= ts.low && index <= ts.high {
		// 从 jumpOffsets 表中查出偏移量进行跳转
		offset = int(ts.jumpOffsets[index - ts.low])
	} else {
		// 使用 defaultOffset 跳转
		offset = int(ts.defaultOffset)
	}
	base.Branch(frame, offset)
}