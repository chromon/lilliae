package classfile

// LineNumberTable 属性表存放方法的行号信息，属于调试信息，都不是运行时必需的
/*
LineNumberTable_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 line_number_table_length;
	{ 	u2 start_pc;
		u2 line_number;
	} line_number_table[line_number_table_length];
}
 */

// 方法行号信息属性表
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPC uint16
	lineNumber uint16
}

// 读取行号信息属性表数据
func (lta *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lta.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lta.lineNumberTable {
		lta.lineNumberTable[i] = &LineNumberTableEntry {
			startPC: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (lta *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lta.lineNumberTable) - 1; i >= 0; i-- {
		entry := lta.lineNumberTable[i]
		if pc >= int(entry.startPC) {
			return int(entry.lineNumber)
		}
	}
	return -1
}