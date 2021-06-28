package heap

import "lilliae/classfile"

// 方法
type Method struct {
	// 从 ClassMember 继承 method 详细信息
	ClassMember
	// 操作数栈大小
	maxStack uint
	// 局部变量表大小
	maxLocals uint
	// 方法字节码
	code []byte
}

// 根据 class 文件中的方法信息创建 Method 表
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

// 复制 class 文件中的方法信息
func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

// 方法是否有 synchronized 关键字
func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags&ACC_SYNCHRONIZED
}

// 方法是否是桥接方法
// 桥接方法（Bridge Method）是一种为了实现某些 Java 语言特性而由编译器自动生成的方法
// 类似于 ACC_SYNTHETIC 合成方法
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags&ACC_BRIDGE
}

// 方法是否在源代码级别采用可变数量的参数
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags&ACC_VARARGS
}

// 是否是本地方法
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags&ACC_NATIVE
}

// 是否是抽象方法
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags&ACC_ABSTRACT
}

// 浮点模式是否采用 FP-strict 模式
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags&ACC_STRICT
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) Code() []byte {
	return m.code
}