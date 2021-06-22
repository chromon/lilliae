package classfile

// 常量数据结构
/*
cp_info {
	u1 tag;
	u1 info[];
}
*/

// 常量是 Java 虚拟机规范严格规定的，共 14 种
// 定义 tag 常量值，用于区分常量类型
const (
	// 表示类或接口的符号引用
	CONSTANT_Class = 7
	// 字段符号引用类型常量
	CONSTANT_Fieldref = 9
	// 非接口方法符号引用类型常量
	CONSTANT_Methodref = 10
	// 接口方法符号引用类型常量
	CONSTANT_InterfaceMethodref = 11
	// 表示 java.lang.String 字面量
	CONSTANT_String = 8
	// 整数类型常量
	CONSTANT_Integer = 3
	// 单精度浮点数类型常量
	CONSTANT_Float = 4
	// 长整形类型常量
	CONSTANT_Long = 5
	// 双精度浮点数类型常量
	CONSTANT_Double = 6
	// 字段或方法的名称和描述符
	// 与 CONSTANT_Class 可以唯一确定一个字段或方法
	CONSTANT_NameAndType = 12
	// MUTF-8 编码字符串类型常量
	CONSTANT_Utf8 = 1
	// 以下三个常量是 JDK 7 之后添加到 class 文件中的，目的是支持 invokedynamic 指令
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 18
)

// 常量数据结构
type ConstantInfo interface {
	// readInfo 方法读取常量信息，需要由具体的常量结构体实现
	readInfo(reader *ClassReader)
}

// 读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// 创建具体的常量对象
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	//case CONSTANT_MethodType:
	//	return &ConstantMethodTypeInfo{}
	//case CONSTANT_MethodHandle:
	//	return &ConstantMethodHandleInfo{}
	//case CONSTANT_InvokeDynamic:
	//	return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag")
	}
}