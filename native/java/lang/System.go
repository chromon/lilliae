package lang

import (
	"lilliae/native"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

const jlSystem = "java/lang/System"

func init() {
	// 注册 arraycopy()
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// 实现 arraycopy()
// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *runtimedataarea.Frame) {
	// 从局部变量表中取出 5 个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	// 源数组和目标数组不能为 null
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	// 源数组和目标数组必须兼容才能拷贝
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// 源数组和目标数组是否兼容
func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	// 源数组和目标数组必须是数组
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}
	// 源数组和目标数组必须是引用数组才可以拷贝
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}