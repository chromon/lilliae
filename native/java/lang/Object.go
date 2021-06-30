package lang

import (
	"lilliae/native"
	"lilliae/runtimedataarea"
	"unsafe"
)

const jlObject = "java/lang/Object"

func init() {
	// 注册 getClass() 本地方法
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

// 实现 getClass() 方法
// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *runtimedataarea.Frame) {
	// 从局部变量表中拿到 this 引用
	this := frame.LocalVars().GetThis()
	// 通过 Class() 方法拿到它的 Class 结构体指针
	// 进而又通过 JClass() 方法拿到它的类对象
	class := this.Class().JClass()
	// 把类对象推入操作数栈顶
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *runtimedataarea.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *runtimedataarea.Frame) {
	this := frame.LocalVars().GetThis()

	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}