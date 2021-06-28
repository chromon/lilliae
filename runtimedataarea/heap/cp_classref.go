package heap

import "lilliae/classfile"

// 类符号引用
type ClassRef struct {
	SymRef
}

// 根据 class 文件中存储的类常量创建 ClassRef 实例
func newClassRef(cp *ConstantPool,
	classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}