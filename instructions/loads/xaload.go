package loads

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// <t>aload 系列指令按索引取数组元素值，然后推入操作数栈

// 从数组中加载引用类型
type AALOAD struct {
	base.NoOperandsInstruction
}

func (a *AALOAD) Execute(frame *Frame) {
	// 从操作数栈中弹出第一个操作数：数组索引，
	// 然后弹出第二个操作数：数组引用
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	// 检查数组引用是否为 null
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	// 检查数组索引是否大于等于 0 且小于数组长度
	checkIndex(len(refs), index)
	// 按索引取出数组元素，推入操作数栈顶
	stack.PushRef(refs[index])
}

// 从数组中加载 byte 或 boolean 类型
type BALOAD struct{
	base.NoOperandsInstruction
}

func (b *BALOAD) Execute(frame *Frame) {
	// 从操作数栈中弹出第一个操作数：数组索引，
	// 然后弹出第二个操作数：数组引用
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	// 检查数组索引是否大于等于 0 且小于数组长度
	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	// 按索引取出数组元素，推入操作数栈顶
	stack.PushInt(int32(bytes[index]))
}

// 从数组中加载 char 类型
type CALOAD struct {
	base.NoOperandsInstruction
}

func (c *CALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// 从数组中加载 double 类型
type DALOAD struct {
	base.NoOperandsInstruction
}

func (d *DALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// 从数组中加载 float 类型
type FALOAD struct {
	base.NoOperandsInstruction
}

func (f *FALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// 从数组中加载 int 类型
type IALOAD struct {
	base.NoOperandsInstruction
}

func (i *IALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// 从数组中加载 long 类型
type LALOAD struct {
	base.NoOperandsInstruction
}

func (l *LALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// 从数组中加载 short 类型
type SALOAD struct {
	base.NoOperandsInstruction
}

func (s *SALOAD) Execute(frame *Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

// 检查数组引用是否为 null
func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

// 如果数组索引小于 0，或者大于等于数组长度，则抛出
// ArrayIndexOutOfBoundsException
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
