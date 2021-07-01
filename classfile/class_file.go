package classfile

import "fmt"

// Java 虚拟机规范描述的 class 文件格式结构体
/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// class 文件结构，包含 Java 虚拟机规范定义的 class 文件格式
type ClassFile struct {
	// u4：魔数，识别 class 文件格式
	magic uint32
	// u2：副版本号
	minorVersion uint16
	// u2：主版本号
	majorVersion uint16
	// 常量池表
	constantPool ConstantPool
	// u2：访问标识
	accessFlags uint16
	// u2：类索引
	thisClass uint16
	// u2：父类索引
	superClass uint16
	// u2：接口索引集合
	interfaces []uint16
	// 字段表
	fields []*MemberInfo
	// 方法表
	methods []*MemberInfo
	// 属性表
	attributes [] AttributeInfo
}

// 将 classData 解析为 ClassFile 结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 依次调用其他方法解析 class 文件
func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

// 读取魔数
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// Java 虚拟机规范对于不符合要求的 class 文件格式会抛出 ClassFormatError 异常
		panic("java.lang.ClassFormatError: magic")
	}
}

// 读取版本号
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	switch cf.majorVersion {
	case 45:
		// JDK 1.02 ~ 1.1
		return
	case 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56 :
		// JDK 1.2 ~ 12
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError")
}

// 获取副版本
func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

// 获取主版本
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

// 获取常量池对象
func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

// 获取访问标记
// 类访问标志是一个 16 位的 bitmask，指出 class 文件定义的是类还是接口，
// 访问级别是 public 还是 private 等
func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

// 字段表，存储字段信息
func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

// 方法表，存储方法信息
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

// 从常量池查找类名
func (cf *ClassFile) ClassName() string {
	// 每个类都有名字，所以类名必须是有效的常量池索引
	return cf.constantPool.getClassName(cf.thisClass)
}

// 从常量池中查找父类名（类似于完全限定名，但是将点换成了斜线）
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		// 除 Object 类之外，其他类都有父类，所以 superClass 只在 Object.class 中是 0
		// 在其他 class 文件中必须是有效的常量池索引
		return cf.constantPool.getClassName(cf.superClass)
	}
	// 只有 java.lang.Object，没有其他父类
	return ""
}

// 从常量池中查找接口名
func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (cf *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range cf.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}