package heap

import "lilliae/classfile"

// 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool,
		refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if imr.method == nil {
		imr.resolveInterfaceMethodRef()
	}
	return imr.method
}

// jvms8 5.4.3.4
func (imr *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := imr.ResolveClass()
	// todo
}