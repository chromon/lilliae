package heap

import "lilliae/classfile"

// 字段
type Field struct {
	// 从 ClassMember 继承 field 详细信息
	ClassMember
	// 常量值索引
	constValueIndex uint
	// 变量槽 id，表示字段在 slots 中的位置
	slotId uint
}

// 根据 class 文件的字段信息创建字段表
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

// 从 class 文件中赋值数据
func (f *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

// 是否有 volatile 关键字
func (f *Field) IsVolatile() bool {
	return 0 != f.accessFlags&ACC_VOLATILE
}

// 是否有 transient 关键字
func (f *Field) IsTransient() bool {
	return 0 != f.accessFlags&ACC_TRANSIENT
}

// 是否是枚举类型
func (f *Field) IsEnum() bool {
	return 0 != f.accessFlags&ACC_ENUM
}

func (f *Field) ConstValueIndex() uint {
	return f.constValueIndex
}

func (f *Field) SlotId() uint {
	return f.slotId
}

// 返回字段是否是 long 或 double 类型
func (f *Field) isLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}