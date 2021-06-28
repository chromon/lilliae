package classfile

// 该常量池为 class 文件中的常量池，应于运行时数据区中 heap 包中的 常量池相区分

// 常量池里面存放着各种的常量信息
// 包括数字和字符串常量、类和接口名、字段和方法名等

// 常量池实际上也是一个表，有三点需要特别注意：
// 1. 表头给出的常量池大小比实际大 1。假设表头给出的值是 n，那么常量池的实际大小是 n – 1
// 2. 有效的常量池索引是 1 ~ n – 1。0 是无效索引，表示不指向任何常量
// 3. CONSTANT_Long_info 和 CONSTANT_Double_info 各占两个位置。即，如果常量池中
// 存在这两种常量，实际的常量数量比 n – 1 还要少，而且 1 ~ n – 1 的某些数也会变成无效索引

// 常量池中的常量分为两类：
// 字面量（literal）：包括数字常量和字符串常量
// 符号引用（symbolic reference）：包括类和接口名、字段和方法信息等
// 除了字面量，其他常量都是通过索引直接或间接指向 CONSTANT_Utf8_info 常量
type ConstantPool []ConstantInfo

// 获取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			// 以上内容占两个位置
			i++
		}
	}
	return cp
}

// 按常量池索引查找常量
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("invalid constant pool index")
}

// 从常量池查找字段或方法的名字和描述符
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池查找类名
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

// 从常量池查找 UTF-8 字符集
func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}