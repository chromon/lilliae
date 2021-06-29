package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// new 指令，专门用于创建类实例
type NEW struct {
	base.Index16Instruction
}

// new 指令的操作数是一个 uint16 索引，来自字节码
func (n *NEW) Execute(frame *runtimedataarea.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 通过索引可以从当前类的运行时常量池中找到一个类符号引用
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	// 解析符号引用，得到类数据
	class := classRef.ResolvedClass()
	// 类的初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 接口和抽象类不能实例化，需要抛出异常
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	// 创建对象
	ref := class.NewObject()
	// 将对象引用推入栈顶
	frame.OperandStack().PushRef(ref)
}