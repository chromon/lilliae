package classfile

// SourceFile 是可选定长属性，只会出现在 ClassFile 结构中，用于指出源文件名
/*
SourceFile_attribute {
	u2 attribute_name_index;
	// attribute_length 的值必须是2
	u4 attribute_length;
	u2 sourcefile_index;
}
 */

// 可选定长属性，用于指出源文件名
type SourceFileAttribute struct {
	cp ConstantPool
	// 常量池索引，指向 CONSTANT_Utf8_info 常量
	sourceFileIndex uint16
}

// 读取属性索引
func (sfa *SourceFileAttribute) readInfo(reader *ClassReader) {
	sfa.sourceFileIndex = reader.readUint16()
}

// 获取文件名
func (sfa *SourceFileAttribute) FileName() string {
	return sfa.cp.getUtf8(sfa.sourceFileIndex)
}