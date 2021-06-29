package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// anewarray 指令用来创建引用类型数组，需要两个操作数，索引和数组长度
type ANEW_ARRAY struct {
	// 索引
	base.Index16Instruction
}

func (a *ANEW_ARRAY) Execute(frame *runtimedataarea.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从当前类的运行时常量池中找到一个类符号引用
	classRef := cp.GetConstant(a.Index).(*heap.ClassRef)
	// 解析符号引用得到数组元素的类（注意：multianewarray 解析得到的是数组类）
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	// 第二个操作数是数组长度，从操作数栈中弹出
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// 返回与类对应的数组类
	arrClass := componentClass.ArrayClass()
	// 创建数组对象并推入操作数栈
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}