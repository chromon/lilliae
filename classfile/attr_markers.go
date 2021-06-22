package classfile

// Deprecated 和 Synthetic 是最简单的两种属性，仅起标记作用，不包含任何数据
/*
Deprecated_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
}
Synthetic_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
}
 */

// 仅用于标记属性的结构体
// 由于不包含任何数据，所以 attribute_length 值为 0
type MarkerAttribute struct {}

// Deprecated 属性用于指出类、接口、字段或方法已经不建议使用
type DeprecatedAttribute struct {
	MarkerAttribute
}

// Synthetic 属性用来标记源文件中不存在、由编译器生成的类成员
// 引入 Synthetic 属性主要是为了支持嵌套类和嵌套接口
type SyntheticAttribute struct {
	MarkerAttribute
}

// 读取属性信息
func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
	// 由于这两个属性都没有数据，所以该方法为空
}