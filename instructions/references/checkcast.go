package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// checkcast 指令和 instanceof 指令很像
// 区别在于：
//	instanceof 指令会改变操作数栈（弹出对象引用，推入判断结果）
//	checkcast 则不改变操作数栈（如果判断失败，直接抛出 ClassCastException异常）
type CHECK_CAST struct {
	base.Index16Instruction
}

func (cc *CHECK_CAST) Execute(frame *runtimedataarea.Frame) {
	// 先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		// 如果引用是 null，则指令执行结束
		// 也就是说，null 引用可以转换成任何类型
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(cc.Index).(*heap.ClassRef)
	// 解析类符号引用
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		// 判断对象是否是类的实例。如果是的话，指令执行结束，否则抛出异常
		panic("java.lang.ClassCastException")
	}
}