package classfile

// LocalVariableTable 属性表中存放方法的局部变量信息，属于调试信息，都不是运行时必需的
/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
// 方法的局部变量信息属性表
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

// 方法的局部变量信息实体类
type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

// 读取部变量信息属性表
func (lvt *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	lvt.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range lvt.localVariableTable {
		lvt.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}