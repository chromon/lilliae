package classfile

// ConstantValue 是定长属性，只会出现在 field_info 结构中，用于表示常量表达式的值
/*
ConstantValue_attribute {
	u2 attribute_name_index;
	// 值必须是 2
	u4 attribute_length;
	u2 constantvalue_index;
}
 */

// 定长属性，用于表示常量表达式的值
type ConstantValueAttribute struct {
	// 常量池索引，但具体指向哪种常量因字段类型而异
	// 如：long 字段类型对应常量类型 CONSTANT_Long_info
	constantValueIndex uint16
}

// 读取常量属性索引值
func (cva *ConstantValueAttribute) readInfo(reader *ClassReader) {
	cva.constantValueIndex = reader.readUint16()
}

// 获取常量表达式的索引值
func (cva *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return cva.constantValueIndex
}