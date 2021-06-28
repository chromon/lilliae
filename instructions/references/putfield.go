package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// putfield 指令给实例变量赋值
type PUT_FIELD struct {
	base.Index16Instruction
}

func (pf *PUT_FIELD) Execute(frame *runtimedataarea.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	// 通过索引从当前类的运行时常量池中查找字段符号引用
	fieldRef := cp.GetConstant(pf.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	// 解析后的字段必须是实例字段，否则抛出异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 如果是 final 字段，则只能在构造函数中初始化，否则抛出 IllegalAccessError
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 先根据字段类型从操作数栈中弹出相应的变量值，然后弹出
	// 对象引用。如果引用是null，需要抛出著名的空指针异常，
	// 否则通过引用给实例变量赋值
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}