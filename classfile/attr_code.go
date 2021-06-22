package classfile

// Code 是变长属性，只存在于 method_info 结构中。Code 属性中存放字节码等方法相关信息
/*
Code_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	// 操作数栈的最大深度
	u2 max_stack;
	// 局部变量表大小
	u2 max_locals;
	u4 code_length;
	u1 code[code_length];
	u2 exception_table_length;
	{ 	u2 start_pc;
		u2 end_pc;
		u2 handler_pc;
		u2 catch_type;
	} exception_table[exception_table_length];
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}
 */

// 变长属性，存放字节码等方法相关信息
type CodeAttribute struct {
	// 常量池
	cp ConstantPool
	// 操作数栈的最大深度
	maxStack uint16
	// 局部变量表大小
	maxLocals uint16
	// 字节码，存在 u1 表中
	code []byte
	// 异常处理表
	exceptionTable []*ExceptionTableEntry
	// 属性表
	attributes []AttributeInfo
}

// 异常处理表
type ExceptionTableEntry struct {
	startPC uint16
	endPC uint16
	handlerPC uint16
	catchType uint16
}

// 读取属性信息
func (ca *CodeAttribute) readInfo(reader *ClassReader) {
	// 操作数栈的最大深度
	ca.maxStack = reader.readUint16()
	// 局部变量表大小
	ca.maxLocals = reader.readUint16()
	// 字节码长度
	codeLength := reader.readUint32()
	// 字节码
	ca.code = reader.readBytes(codeLength)
	// 异常表
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader, ca.cp)
}

// 读取异常表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPC:   reader.readUint16(),
			endPC:     reader.readUint16(),
			handlerPC: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}