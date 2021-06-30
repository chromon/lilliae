package reserved

import (
	"lilliae/instructions/base"
	"lilliae/native"
	"lilliae/runtimedataarea"
	// 如果没有任何包依赖 lang 包，它就不会被编译进可执行文件，
	// 本地方法也就不会被注册。所以需要一个地方导入 lang 包
	// 由于没有显示使用 lang 中的变量或函数，所以必须在包名前面加上下划线
	_ "lilliae/native/java/lang"
	_ "lilliae/native/sun/misc"
)

// 0xFE 指令

// 调用本地方法
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (n *INVOKE_NATIVE) Execute(frame *runtimedataarea.Frame) {

	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	// 根据类名、方法名和方法描述符从本地方法注册表中查找本地方法实现
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		// 找不到则抛出异常
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	// 直接调用本地方法
	nativeMethod(frame)
}