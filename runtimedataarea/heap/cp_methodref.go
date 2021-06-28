package heap

import "lilliae/classfile"

// 方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool,
		refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (mr *MethodRef) ResolvedMethod() *Method {
	if mr.method == nil {
		mr.resolveMethodRef()
	}
	return mr.method
}

// jvms8 5.4.3.3
func (mr *MethodRef) resolveMethodRef() {
	//class := mr.Class()
	// todo
}