package heap

import "lilliae/classfile"

// 非接口方法的符号引用

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

// 解析方法符号引用
func (mr *MethodRef) ResolvedMethod() *Method {
	if mr.method == nil {
		// 符号引用还没有被解析过，调用相关方法进行解析
		mr.resolveMethodRef()
	}
	return mr.method
}

// 解析非接口方法符号引用
func (mr *MethodRef) resolveMethodRef() {
	// 类 d 想通过方法符号引用访问类 c 的某个方法
	d := mr.cp.class
	// 先解析符号引用得到类 c
	c := mr.ResolvedClass()

	// 如果 c 是接口则抛出异常
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据方法名和描述符查找方法
	method := lookupMethod(c, mr.name, mr.descriptor)
	if method == nil {
		// 查找不到对应的方法则抛出异常
		panic("java.lang.NoSuchMethodError")
	}

	// 检查类 d 是否有权限访问该方法
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	mr.method = method
}

// 根据方法名和描述符查找方法
func lookupMethod(class *Class, name, descriptor string) *Method {
	// 从 c 的继承层次中查找
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		// 如果继承层次中找不到就到接口中去找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}