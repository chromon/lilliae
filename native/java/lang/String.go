package lang

import (
	"lilliae/native"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *runtimedataarea.Frame) {
	this := frame.LocalVars().GetThis()
	// 检查字符串常量池中是否有当前字符串，如果没有则放入并返回该字符串，否则找到并直接返回
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}