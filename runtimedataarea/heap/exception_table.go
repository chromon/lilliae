package heap

import "lilliae/classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPC   int
	endPC     int
	handlerPC int
	catchType *ClassRef
}

// class文件中的异常处理表转换成 ExceptionTable 类型
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPC:   int(entry.StartPC()),
			endPC:     int(entry.EndPc()),
			handlerPC: int(entry.HandlerPC()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

// 从运行时常量池中查找类符号引用
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	// 异常处理项的 catchType 有可能是 0
	// 0 是无效的常量池索引，但是在这里 0 并非表示 catch-none，而是表示 catch-all
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 查找异常处理表
func (t ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range t {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPC && pc < handler.endPC {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}