package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// multianewarray 指令创建多维数组
type MULTI_ANEW_ARRAY struct {
	// 索引值，通过这个索引可以从运行时常量池中找到一个类符号引用，
	// 解析这个引用就可以得到多维数组类
	index uint16
	// 表示数组维度
	dimensions uint8
}

// 读取两个操作数，这两个操作数在字节码中紧跟在指令操作码后面
func (a *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	a.index = reader.ReadUint16()
	a.dimensions = reader.ReadUint8()
}

func (a *MULTI_ANEW_ARRAY) Execute(frame *runtimedataarea.Frame) {
	// 通过索引从运行时常量池中找到类符号引用，解析引用得到多维数组类
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(a.index)).(*heap.ClassRef)
	// 在 anewarray 指令中，解析类符号引用后得到的是数组元素的类，
	// 而这里解析出来的直接就是数组类
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	// 从操作数栈中弹出 n 个整数，分别代表每一个维度的数组长度
	counts := popAndCheckCounts(stack, int(a.dimensions))
	// 根据数组类、数组维度和各个维度的数组长度创建多维数组
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

// 从操作数栈中弹出n个整数，分别代表每一个维度的数组长度
func popAndCheckCounts(stack *runtimedataarea.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	// 从操作数栈中弹出 n 个 int 值，并且确保它们都大于等于 0。
	// 如果其中任何一个小于 0，则抛出异常
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

// 根据数组类、数组维度和各个维度的数组长度创建多维数组
func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}