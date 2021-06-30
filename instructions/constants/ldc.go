package constants

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// ldc 系列指令从运行时常量池中加载常量值，并把它推入操作数栈。
// ldc 系列指令属于常量类指令，共 3 条。
// 其中 ldc 和 ldc_w 指令用于加载 int、float 和字符串常量，
// java.lang.Class 实例或者 MethodType 和 MethodHandle 实例。
// ldc2_w 指令用于加载 long 和 double 常量。ldc 和 ldc_w 指令的区别仅在于操作数的宽度

type LDC struct {
	base.Index8Instruction
}

func (l *LDC) Execute(frame *runtimedataarea.Frame) {
	_ldc(frame, l.Index)
}

// Push item from run-time constant pool (wide index)
type LDC_W struct {
	base.Index16Instruction
}

func (lw *LDC_W) Execute(frame *runtimedataarea.Frame) {
	_ldc(frame, lw.Index)
}

func _ldc(frame *runtimedataarea.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	case string:
		// 如果 ldc 试图从运行时常量池中加载字符串常量，则先通过常量拿到 Go 字符串，
		// 然后把它转成 Java 字符串实例并把引用推入操作数栈顶
		internedStr := heap.JString(class.Loader(), c.(string))
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		// 加载类对象
		// 如果运行时，常量池中的常量是类引用，则解析类引用，
		// 然后把类的类对象推入操作数栈顶
		classRef := c.(*heap.ClassRef)
		classObj := classRef.ResolvedClass().JClass()
		stack.PushRef(classObj)
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct {
	base.Index16Instruction
}

func (l *LDC2_W) Execute(frame *runtimedataarea.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(l.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}