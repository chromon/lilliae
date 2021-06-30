package native

import "lilliae/runtimedataarea"

// 本地方法类型定义为一个函数，参数是 Frame 结构体，frame 是本地方法的工作空间
type NativeMethod func(frame *runtimedataarea.Frame)

// 本地方法注册表，用来注册和查找本地方法
var registry = map[string]NativeMethod{}

// 注册本地方法
func Register(className, methodName, methodDescriptor string,
		method NativeMethod) {
	// 类名、方法名和方法描述符加在一起才能唯一确定一个方法，
	// 所以把它们的组合作为本地方法注册表的键
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 根据类名、方法名和方法描述符查找本地方法实现
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	// java.lang.Object等类是通过一个 registerNatives 的本地方法来注册其他本地方法的
	// 而将使用自己注册所有的本地方法实现，所以 registerNatives 方法就没有用处了
	// 直接返回一个空的实现
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

// 空的本地方法实现
func emptyNativeMethod(frame *runtimedataarea.Frame) {
	// do nothing
}