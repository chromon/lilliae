package control

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// Java 语言中的 switch-case 语句有两种实现方式：
// 如果 case 值可以编码成一个索引表，则实现成 tableswitch 指令；
// 否则实现成 lookupswitch 指令

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/

type LOOKUP_SWITCH struct {
	// 对应默认情况下执行跳转所需的字节码偏移量
	defaultOffset int32
	//
	npairs int32
	// 类似于 Map，key 是 case 值，value 是跳转偏移量
	matchOffsets []int32
}

func (ls *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// 指令操作码的后面有 0 ~ 3 字节的 padding，
	// 以保证 defaultOffset 在字节码中的地址是 4 的倍数
	reader.SkipPadding()
	ls.defaultOffset = reader.ReadInt32()
	ls.npairs = reader.ReadInt32()
	ls.matchOffsets = reader.ReadInt32s(ls.npairs * 2)
}

func (ls *LOOKUP_SWITCH) Execute(frame *runtimedataarea.Frame) {
	// 操作数栈中弹出一个 int变量
	key := frame.OperandStack().PopInt()
	// 由 key 查找 matchOffsets，看是否能找到匹配的 key
	for i := int32(0); i < ls.npairs * 2; i += 2 {
		if ls.matchOffsets[i] == key {
			// 找到匹配的 key，按照 value 给出的偏移量跳转
			offset := ls.matchOffsets[i + 1]
			base.Branch(frame, int(offset))
			return
		}
	}
	// 否则使用默认的 defaultOffset 跳转
	base.Branch(frame, int(ls.defaultOffset))
}