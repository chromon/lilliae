package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// getfield 指令获取对象的实例变量值，然后推入操作数栈
type GET_FIELD struct {
	base.Index16Instruction
}

func (gf *GET_FIELD) Execute(frame *runtimedataarea.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从当前类的运行时常量池中查找字段符号引用
	fieldRef := cp.GetConstant(gf.Index).(*heap.FieldRef)
	// 解析
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 弹出对象引用，如果是 null，则抛出 NullPointerException
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 根据字段类型，获取相应的实例变量值，然后推入操作数栈
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()

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
