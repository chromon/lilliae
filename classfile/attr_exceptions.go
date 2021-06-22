package classfile

// Exceptions 是变长属性，记录方法抛出的异常表
/*
Exceptions_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 number_of_exceptions;
	u2 exception_index_table[number_of_exceptions];
}
 */

// 异常表
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

// 读取异常表
func (ea *ExceptionsAttribute) readInfo(reader *ClassReader) {
	ea.exceptionIndexTable = reader.readUint16s()
}

// 获取异常表
func (ea *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return ea.exceptionIndexTable
}