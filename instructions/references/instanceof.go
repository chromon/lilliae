package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// 判断对象是否是某个类的实例（或者对象的类是否实现了某个接口），并将结果推入操作数栈
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (ins *INSTANCE_OF) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	// 对象引用
	ref := stack.PopRef()
	if ref == nil {
		// 如果是 null，则把 0 推入操作数栈（即 instanceof 判断结果为 false）
		stack.PushInt(0)
		return
	}

	// 通过索引从当前类的运行时常量池中查找类符号引用
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(ins.Index).(*heap.ClassRef)
	// 解析类符号引用
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}