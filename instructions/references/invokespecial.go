package references

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// invokespecial 指令用来调用无须动态绑定的实例方法，
// 包括构造函数、私有方法和通过 super 关键字调用的超类方法

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (s *INVOKE_SPECIAL) Execute(frame *runtimedataarea.Frame) {
	// 当前类
	currentClass := frame.Method().Class()
	// 当前常量池
	cp := currentClass.ConstantPool()
	// 方法符号引用
	methodRef := cp.GetConstant(s.Index).(*heap.MethodRef)
	// 解析符号引用得到类和方法
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	// 解析出的方法是构造方法时，声明该方法的类必须是解析出来的类
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 如果解析出来的方法是静态方法则抛出异常
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中弹出 this 引用
	// GetRefFromTop 返回距离操作数栈顶 n 个单元格的引用变量
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		// 如果为 nil 则抛出异常
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

	// 如果调用的超类中的函数，但不是构造函数，且当前类的 ACC_SUPER 标志被设置，
	// 需要一个额外的过程查找最终要调用的方法；
	// 否则前面从方法符号引用中解析出来的方法就是要调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
			resolvedClass.IsSuperClassOf(currentClass) &&
			resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}

	// 如果查找过程失败，或者找到的方法是抽象的，抛出 AbstractMethodError 异常
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	// 调用方法
	base.InvokeMethod(frame, methodToBeInvoked)
}