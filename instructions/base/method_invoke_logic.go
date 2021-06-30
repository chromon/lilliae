package base

import (
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// Java 虚拟机规范一共提供了 4 条方法调用指令
// invokestatic、invokespecial、invokeinterface、invokevirtual
// 四条方法调用指令内容基本相同：
// 定位到需要调用的方法之后，Java 虚拟机要给这个方法创建
// 一个新的帧并把它推入 Java 虚拟机栈顶，然后传递参数

// 方法调用指令
func InvokeMethod(invokerFrame *runtimedataarea.Frame, method *heap.Method) {
	// 创建新的栈帧并推入 Java 虚拟机栈
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	// 方法参数在局部变量表中占用 slot 数量
	// 不一定是局部变量表中的参数个数，double 和 long 占两个
	// 另外对于实例方法，Java 编译器会在参数列表的前面添加一个参数 this 引用
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}