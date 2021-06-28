package heap

import (
	"fmt"
	"lilliae/classfile"
)

// 运行时数据区常量池

// 运行时常量池主要存放两类信息：字面量（literal）和符号引用（symbolic reference）
// 字面量包括整数、浮点数和字符串字面量
// 符号引用包括类符号引用、字段符号引用、方法符号引用和接口方法符号引用

// 常量
type Constant interface {}

// 常量池
type ConstantPool struct {
	// 类指针
	class *Class
	// 常量
	consts []Constant
}

// 将 class 文件中的常量池转换成运行时常量池
// 核心逻辑将 []classfile.ConstantInfo 转换成 []heap.Constant
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	// class 文件中常量池大小
	cpCount := len(cfCp)
	// 构建运行时常量池数据数组
	consts := make([]Constant, cpCount)
	// 构建运行时常量池
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			// int 型常量，直接取出，放入 consts 数组
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			// float 型常量，直接取出，放入 consts 数组
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.ConstantLongInfo:
			// long 型常量，直接取出，放入 consts 数组
			// long 型常量在常量池中占 2 个位置，索引 i 需要额外 + 1
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			// double 型常量，直接取出，放入 consts 数组
			// double 型常量在常量池中占 2 个位置，索引 i 需要额外 + 1
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			// string 型常量，直接取出，放入 consts 数组
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.ConstantClassInfo:
			// 类符号引用
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			// 字段符号引用
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			// 方法符号引用
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			// 接口方法符号引用
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:
			// todo
		}
	}
	return rtCp
}

// 根据索引返回常量
func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}