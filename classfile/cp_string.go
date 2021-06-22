package classfile

/*
CONSTANT_String_info {
	u1 tag;
	u2 string_index;
}
 */

// CONSTANT_String_info 常量表示 java.lang.String 字面量
// CONSTANT_String_info本身并不存放字符串数据，只存放常量池索引，
// 这个索引指向一个 CONSTANT_Utf8_info 常量
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

// 读取常量池索引
func (csi *ConstantStringInfo) readInfo(reader *ClassReader) {
	csi.stringIndex = reader.readUint16()
}

// 按照索引从常量池中查找字符串
func (csi *ConstantStringInfo) String() string {
	return csi.cp.getUtf8(csi.stringIndex)
}