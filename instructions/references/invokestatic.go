package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// invokestatic 指令，调用静态方法

// 调用类的静态方法
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (s *INVOKE_STATIC) Execute(frame *runtimedataarea.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(s.Index).(*heap.MethodRef)
	// 解析方法
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		// 必须是静态方法
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	// 不能是类初始化方法
	// 类初始化方法只能由 Java 虚拟机调用，不能使用 invokestatic 指令调用
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}