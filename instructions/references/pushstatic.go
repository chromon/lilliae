package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// pushstatic 和 getstatic 指令用于存取静态变量

// 给类的静态变量赋值
type PUT_STATIC struct {
	base.Index16Instruction
}

func (ps *PUT_STATIC) Execute(frame *runtimedataarea.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 通过索引从当前类的运行时常量池中查找字段符号引用
	fieldRef := cp.GetConstant(ps.Index).(*heap.FieldRef)
	// 解析符号引用，得到类的静态变量
	field := fieldRef.ResolvedField()
	class := field.Class()
	// 类的初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 解析后字段是实例字段而非静态字段，则抛出异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是 final 字段，则实际操作的是静态常量，只能在类初始化方法中给它赋值
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 根据字段类型从操作数栈中弹出相应的值，赋给静态变量
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}