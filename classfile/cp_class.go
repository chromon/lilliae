package classfile

/*
CONSTANT_Class_info {
	u1 tag;
	u2 name_index;
}
 */

// CONSTANT_Class_info 常量表示类或者接口的符号引用
type ConstantClassInfo struct {
	// 常量池
	cp ConstantPool
	// 常量池索引
	nameIndex uint16
}

// 读取常量池索引
func (cci *ConstantClassInfo) readInfo(reader *ClassReader) {
	cci.nameIndex = reader.readUint16()
}

// 按照索引从常量池中查找字符串
func (cci *ConstantClassInfo) Name() string {
	return cci.cp.getUtf8(cci.nameIndex)
}