package heap

import "lilliae/classfile"

// 字段符号引用
type FieldRef struct {
	MemberRef
	// 缓存解析后的字段指针
	field *Field
}

// 创建 FieldRef 实例
func newFieldRef(cp *ConstantPool,
		refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 解析字段符号引用
func (fr *FieldRef) ResolvedField() *Field {
	if fr.field == nil {
		fr.resolveFieldRef()
	}
	return fr.field
}

// 根据 Java 虚拟机规范 5.4.3.2 节给出了字段符号引用的解析步骤
// 如果类 D 想通过字段符号引用访问类 C 的某个字段，
// 首先要解析符号引用得到类 C，然后根据字段名和描述符查找字段。
// 如果字段查找失败，则虚拟机抛出 NoSuchFieldError 异常。
// 如果查找成功，但 D 没有足够的权限访问该字段，则虚拟机抛出 IllegalAccessError 异常
func (fr *FieldRef) resolveFieldRef() {
	d := fr.cp.class
	c := fr.ResolvedClass()
	field := lookupField(c, fr.name, fr.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	fr.field = field
}

// 字段查找
// 首先在 C 的字段中查找。如果找不到，在 C 的直接接口递归应用这个查找过程。
// 如果还找不到的话，在 C 的超类中递归应用这个查找过程。
// 如果仍然找不到，则查找失败
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}