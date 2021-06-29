package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// getstatic 指令和 putstatic 正好相反，
// getstatic 取出类的某个静态变量值，然后推入栈顶
type GET_STATIC struct {
	base.Index16Instruction
}

func (gs *GET_STATIC) Execute(frame *runtimedataarea.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从当前类的运行时常量池中查找字段符号引用
	fieldRef := cp.GetConstant(gs.Index).(*heap.FieldRef)
	// 解析符号引用，得到类的静态变量
	field := fieldRef.ResolvedField()
	class := field.Class()
	// 类的初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 如果解析后的字段不是静态字段则抛出异常
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据字段类型，从静态变量中取出相应的值，然后推入操作数栈顶
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}