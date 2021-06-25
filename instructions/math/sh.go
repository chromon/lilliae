package math

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// 位移指令，分为左移和右移两种，右移指令又可以分为算术右移（有符号右移）
// 和逻辑右移（无符号右移）两种。算术右移和逻辑位移的区别仅在于符号位的扩展
// int x = -1;
// println(Integer.toBinaryString(x));      // 11111111111111111111111111111111
// println(Integer.toBinaryString(x >> 8)); // 11111111111111111111111111111111
// println(Integer.toBinaryString(x >>> 8));// 00000000111111111111111111111111

// int 左位移
type ISHL struct {
	base.NoOperandsInstruction
}

func (isl *ISHL) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	// 要移的比特位数
	v2 := stack.PopInt()
	// 要进行位移的操作数
	v1 := stack.PopInt()
	// 因为 int 变量只有 32 位，所以 v2 只取前五个比特就足够表示位移位数了
	// 与 31（0x1f） 做与运算，得出移位位数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// int 右移位（算术右移）
type ISHR struct {
	base.NoOperandsInstruction
}

func (isr ISHR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	// 要移的比特位数
	v2 := stack.PopInt()
	// 将要移位的操作数
	v1 := stack.PopLong()
	// long 类型变量有 64 位，取 v2 前 6 个比特
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// int 无符号右移位（逻辑右移）
type IUSHR struct {
	base.NoOperandsInstruction
}

func (ius *IUSHR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	// GO 语言没有 >>> 运算符，为了实现无符号右移
	// 先将 v1 转成无符号正数，位移之后，在转回有符号正数
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// long 左位移
type LSHL struct {
	base.NoOperandsInstruction
}

func (lsl *LSHL) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// long 右移位（算术右移）
type LSHR struct {
	base.NoOperandsInstruction
}

func (lsr *LSHR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// long 无符号右移位（逻辑右移）
type LUSHR struct {
	base.NoOperandsInstruction
}

func (lus *LUSHR) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}