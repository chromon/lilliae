package classfile

// 字段表和方法表，分别存储字段和方法信息。字段和方法的基本结构大致相同，差别仅在于属性表
// 下面是 Java 虚拟机规范给出的字段和方法结构定义
/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// 字段和方法表
type MemberInfo struct {
	// 常量池对象
	cp ConstantPool
	// 访问标志（常量池索引）
	accessFlags uint16
	// 字段名或方法名（常量池索引）
	nameIndex uint16
	// 字段或方法描述符（常量池索引）
	descriptorIndex uint16
	// 属性表
	attributes []AttributeInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([] *MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo {
		cp: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

// 获取访问标记
func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

// 获取字段或方法名
func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

// 获取字段或方法描述
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

// 获取 Code 属性
func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		// 类型断言
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}