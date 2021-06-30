package references

import (
	"fmt"
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// 当 Java 虚拟机通过 invokevirtual 调用方法时，
// this 引用指向某个类（或其子类）的实例。因为类的继承层次是固定的，
// 所以虚拟机可以使用一种叫作 vtable（Virtual Method Table）的技术加速方法查找

// 动态绑定
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (v *INVOKE_VIRTUAL) Execute(frame *runtimedataarea.Frame) {
	// 当前类
	currentClass := frame.Method().Class()
	// 当前常量池
	cp := currentClass.ConstantPool()
	// 方法符号引用
	methodRef := cp.GetConstant(v.Index).(*heap.MethodRef)
	// 解析符号引用得到方法
	resolvedMethod := methodRef.ResolvedMethod()
	// 如果解析出来的方法是静态方法则抛出异常
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中弹出 this 引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	// 如果为 nil 则抛出异常
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
	}

	// 确保 protected 方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() &&
			resolvedMethod.Class().IsSuperClassOf(currentClass) &&
			resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
			ref.Class() != currentClass &&
			!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	// 从对象的类中查找到真正要调用的方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(),
		methodRef.Name(), methodRef.Descriptor())
	// 如果找不到或找到的是抽象方法则抛出异常
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

// hack!
func _println(stack *runtimedataarea.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}