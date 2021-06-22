package classfile

// Java虚拟机规范使用属性名来区别不同的属性
// 属性数据放在属性名之后的 u1 表中，这样 Java 虚拟机实现就可以跳过无法识别的属性

/*
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
 */

// 属性数据结构
type AttributeInfo interface {
	// readInfo 方法读取属性信息，需要由具体的属性结构体实现
	readInfo(reader *ClassReader)
}

// 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// 读取单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 读取属性名索引
	attrNameIndex := reader.readUint16()
	// 由属性名索引在常量池中查找属性名
	attrName := cp.getUtf8(attrNameIndex)
	// 读取属性长度
	attrLen := reader.readUint32()
	// 构建属性实例
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// 创建具体属性实例
// Java 虚拟机规范预定义了 23 种属性
// 按照用途，23 种预定义属性可以分为三组：
// 第一组属性是实现 Java 虚拟机所必需的，共有 5 种
// 第二组属性是 Java 类库所必需的共有 12 种
// 第三组属性主要提供给工具使用，共有 6 种。第三组属性是可选的，也就是说可以不出现在 class 文件中
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		// 变长属性，只存在于 method_info 结构中，用于存放字节码等方法相关信息
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		// 定长属性，只会出现在 field_info 中，用于表示常量表达式的值
		return &ConstantValueAttribute{}
	case "Deprecated":
		// 属性用于指出类、接口、字段或方法已不建议使用，仅起到标记作用，不包含任何数据
		return &DeprecatedAttribute{}
	case "Exceptions":
		// 变长属性，记录方法抛出的异常表
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		// 异常堆栈中显示方法行号
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		// 该属性表中存放方法的局部变量信息
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		// 可选定长属性，只会出现在 ClassFile 结构中，用于指出源文件名
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		// 属性用来标记源文件中不存在、由编译器生成的类成员
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}