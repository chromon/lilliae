package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// newarray 指令用来创建基本类型数组，包括 boolean[]、byte[]、
// char[]、short[]、int[]、long[]、float[] 和 double[] 8 种

// newarray 指令需要两个操作数。第一个操作数是一个 uint8 整数，
// 在字节码中紧跟在指令操作码后面，表示要创建哪种类型的数组。
// Java 虚拟机规范把这个操作数叫作 atype，并且规定了它的有效值
const (
	// 数组类型 atype
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// 创建基本类型数组
type NEW_ARRAY struct {
	atype uint8
}

// 读取 atype 值
func (a *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	a.atype = reader.ReadUint8()
}

func (a *NEW_ARRAY) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	// newarray 指令的第二个操作数是 count，从操作数栈中弹出，表示数组长度
	count := stack.PopInt()
	// 如果数组长度小于 0 抛异常
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 根据 atype 值使用当前类的类加载器加载数组类，
	// 然后创建数组对象并推入操作数栈
	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, a.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

// 根据 atype 值使用当前类的类加载器加载数组类
func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}