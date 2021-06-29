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

// 解析接口方法的符号引用
func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if imr.method == nil {
		imr.resolveInterfaceMethodRef()
	}
	return imr.method
}

// 解析接口方法的符号引用具体实现
func (imr *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := imr.cp.class
	// 先解析符号引用得到接口 c
	c := imr.ResolvedClass()
	// 如果 c 不是接口抛出异常
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据方法名和描述符查找方法
	method := lookupInterfaceMethod(c, imr.name, imr.descriptor)
	if method == nil {
		// 查找不到对应的方法则抛出异常
		panic("java.lang.NoSuchMethodError")
	}
	// 检查类 d 是否有权限访问该方法
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	imr.method = method
}

// 根据方法名和描述符查找方法
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		// 如果能在接口中找到方法，直接返回方法
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	// 接口中查找不到方法就在父接口中查找
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}