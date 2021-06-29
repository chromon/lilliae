package heap

import "lilliae/classfile"

// 字段和方法中相同的信息（访问标志、名字、描述符）
type ClassMember struct {
	// 访问标志
	accessFlags uint16
	// 名字
	name string
	// 描述符
	descriptor string
	// 类指针，可以通过字段或方法访问到所属类
	class *Class
}

// 从 class 文件中复制数据
func (cm *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	cm.accessFlags = memberInfo.AccessFlags()
	cm.name = memberInfo.Name()
	cm.descriptor = memberInfo.Descriptor()
}

// 是否是 public
func (cm *ClassMember) IsPublic() bool {
	return 0 != cm.accessFlags&ACC_PUBLIC
}

// 是否是 private
func (cm *ClassMember) IsPrivate() bool {
	return 0 != cm.accessFlags&ACC_PRIVATE
}

// 是否是 protected
func (cm *ClassMember) IsProtected() bool {
	return 0 != cm.accessFlags&ACC_PROTECTED
}

// 是否是 static
func (cm *ClassMember) IsStatic() bool {
	return 0 != cm.accessFlags&ACC_STATIC
}

// 是否是 final
func (cm *ClassMember) IsFinal() bool {
	return 0 != cm.accessFlags&ACC_FINAL
}

// 是否是用户代码产生的
func (cm *ClassMember) IsSynthetic() bool {
	return 0 != cm.accessFlags&ACC_SYNTHETIC
}

func (cm *ClassMember) Name() string {
	return cm.name
}

func (cm *ClassMember) Descriptor() string {
	return cm.descriptor
}

func (cm *ClassMember) Class() *Class {
	return cm.class
}

// 是否是可访问的 public 或 同一个包中，或不同包的子类
func (cm *ClassMember) isAccessibleTo(d *Class) bool {
	// 如果字段是public，则任何类都可以访问
	if cm.IsPublic() {
		return true
	}
	c := cm.class
	// 如果字段是 protected，则只有子类和同一个包下的类可以访问
	if cm.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	// 如果字段有默认访问权限（非 public，非 protected，也非 privated），
	// 则只有同一个包下的类可以访问
	if !cm.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}