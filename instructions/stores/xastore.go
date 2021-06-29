package stores

import (
	"lilliae/instructions/base"
	. "lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// <t>astore 系列指令按索引给数组元素赋值

// 按索引给 引用 数组赋值
type AASTORE struct {
	base.NoOperandsInstruction
}

func (a *AASTORE) Execute(frame *Frame) {
	// aastore 指令的三个操作数分别是：要赋给数组元素的值、数组
	// 索引、数组引用，依次从操作数栈中弹出
	stack := frame.OperandStack()
	ref := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	// 检查数组引用是否为 null
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	// 数组索引应大于等于 0 小于数组长度
	checkIndex(len(refs), index)
	// 按索引给数组元素赋值
	refs[index] = ref
}

// 按索引给 byte 或 boolean 数组赋值
type BASTORE struct {
	base.NoOperandsInstruction
}

func (b *BASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(val)
}

// 按索引给 char 数组赋值
type CASTORE struct {
	base.NoOperandsInstruction
}

func (c *CASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	chars[index] = uint16(val)
}

// 按索引给 double 数组赋值
type DASTORE struct {
	base.NoOperandsInstruction
}

func (d *DASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	doubles[index] = float64(val)
}

// 按索引给 float 数组赋值
type FASTORE struct {
	base.NoOperandsInstruction
}

func (f *FASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	floats[index] = float32(val)
}

// 按索引给 int 数组赋值
type IASTORE struct {
	base.NoOperandsInstruction
}

func (i *IASTORE) Execute(frame *Frame) {
	// iastore 指令的三个操作数分别是：要赋给数组元素的值、数组
	// 索引、数组引用，依次从操作数栈中弹出
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	// 检查数组引用是否为 null
	checkNotNil(arrRef)
	ints := arrRef.Ints()
	// 数组索引应大于等于 0 小于数组长度
	checkIndex(len(ints), index)
	// 按索引给数组元素赋值
	ints[index] = int32(val)
}

// 按索引给 long 数组赋值
type LASTORE struct {
	base.NoOperandsInstruction
}

func (l *LASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	longs[index] = int64(val)
}

// 按索引给 short 数组赋值
type SASTORE struct {
	base.NoOperandsInstruction
}

func (s *SASTORE) Execute(frame *Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(val)
}

// 检查数组引用是否为 null
func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

// 数组索引应大于等于 0 小于数组长度
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}