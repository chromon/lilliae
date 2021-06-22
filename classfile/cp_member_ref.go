package classfile

// CONSTANT_Fieldref_info 表示字段符号引用
// CONSTANT_Methodref_info 表示普通（非接口）方法符号引用
// CONSTANT_InterfaceMethodref_info 表示接口方法符号引用

// 以上三种结构完全相同，例如 CONSTANT_Fieldref_info 如下
/*
CONSTANT_Fieldref_info {
	u1 tag;
	u2 class_index;
	u2 name_and_type_index;
}
 */

// 定义统一的结构体表示以上三种常量
type ConstantMemberrefInfo struct {
	// 常量池
	cp ConstantPool
	// 常量池索引指向 CONSTANT_Class_info
	classIndex uint16
	// 常量池索引指向 CONSTANT_NameAndType_info
	nameAndTypeIndex uint16
}

// 分别从常量池中读取索引数据
func (cmi *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cmi.classIndex = reader.readUint16()
	cmi.nameAndTypeIndex = reader.readUint16()
}

// 由索引查询数据
func (cmi *ConstantMemberrefInfo) ClassName() string {
	return cmi.cp.getClassName(cmi.classIndex)
}

// 有索引查询数据
func (cmi *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return cmi.cp.getNameAndType(cmi.nameAndTypeIndex)
}

// 字段符号引用继承自 ConstantMemberrefInfo，GO 中通过结构体嵌套模拟继承
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

// 普通（非接口）方法符号引用
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

// 接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}