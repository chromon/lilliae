package lang

import (
	"lilliae/native"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

const jlClass = "java/lang/Class"

func init() {
	native.Register(jlClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(jlClass, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(jlClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	//native.Register(jlClass, "isInterface", "()Z", isInterface)
	//native.Register(jlClass, "isPrimitive", "()Z", isPrimitive)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *runtimedataarea.Frame) {
	// 从局部变量表中拿到类名（Java 字符串）
	nameObj := frame.LocalVars().GetRef(0)
	// 转成 Go 字符串
	name := heap.GoString(nameObj)

	// 调用类加载器的LoadClass() 方法获取基本类型的类
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *runtimedataarea.Frame) {
	// 从局部变量表中拿到 this 引用，这是一个类对象引用
	this := frame.LocalVars().GetThis()
	// 获得与之对应的 Class 结构体指针
	class := this.Extra().(*heap.Class)

	// 获取类名
	name := class.JavaName()
	// 转换成 Java 字符串（格式：例如 java.lang.Object）
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// 断言
// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *runtimedataarea.Frame) {
	// todo
	frame.OperandStack().PushBoolean(false)
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *runtimedataarea.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *runtimedataarea.Frame) {
	vars := frame.LocalVars()
	this := vars.GetThis()
	class := this.Extra().(*heap.Class)

	stack := frame.OperandStack()
	stack.PushBoolean(class.IsPrimitive())
}