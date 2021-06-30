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
	// 方法参数在局部变量表中占用 slot 数量
	argSlotCount uint
}

// 根据 class 文件中的方法信息创建 Method 表
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		// 如果是本地方法注入字节码和其他信息
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

// 复制 class 文件中的方法信息
func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

// 本地方法注入字节码和其他信息
func (m *Method) injectCodeAttribute(returnType string) {
	// 本地方法在 class 文件中没有 Code 属性，所以需要给 maxStack 和 maxLocals 字段赋值
	// 本地方法帧的操作数栈至少要能容纳返回值，暂时给 maxStack 字段赋值为 4
	// 因为本地方法帧的局部变量表只用来存放参数值，所以把 argSlotCount 赋给 maxLocals 字段
	m.maxStack = 4 // todo
	m.maxLocals = m.argSlotCount
	// 本地方法的字节码 code 字段，第一条指令都是 0xFE，
	// 第二条指令则根据函数的返回值选择相应的返回指令
	switch returnType[0] {
	case 'V':
		m.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		m.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		m.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		m.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		m.code = []byte{0xfe, 0xad} // lreturn
	default:
		m.code = []byte{0xfe, 0xac} // ireturn
	}
}

// 计算方法参数数量
func (m *Method) calcArgSlotCount(paramTypes []string) {
	// 分解方法描述符，返回 MethodDescriptor 结构体
	for _, paramType := range paramTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount++ // `this` reference
	}
}

// 方法是否有 synchronized 关键字
func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags & ACC_SYNCHRONIZED
}

// 方法是否是桥接方法
// 桥接方法（Bridge Method）是一种为了实现某些 Java 语言特性而由编译器自动生成的方法
// 类似于 ACC_SYNTHETIC 合成方法
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags & ACC_BRIDGE
}

// 方法是否在源代码级别采用可变数量的参数
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags & ACC_VARARGS
}

// 是否是本地方法
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags & ACC_NATIVE
}

// 是否是抽象方法
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags & ACC_ABSTRACT
}

// 浮点模式是否采用 FP-strict 模式
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags & ACC_STRICT
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

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}