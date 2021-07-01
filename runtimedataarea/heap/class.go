package heap

import (
	"lilliae/classfile"
	"strings"
)

// 将要放入方法区内的类
type Class struct {
	// 类的访问标志
	accessFlags uint16
	// 类名，完全限定名如 java/lang/Object
	name string
	// 父类名，完全限定名
	superClassName string
	// 接口名，完全限定名
	interfaceNames []string
	// 运行时常量池指针
	constantPool *ConstantPool
	// 字段表
	fields []*Field
	// 方法表
	methods []*Method
	// 类加载器
	loader *ClassLoader
	// 父类指针
	superClass *Class
	// 借口指针
	interfaces []*Class
	// 实例变量占据的空间大小
	instanceSlotCount uint
	// 类变量占据空间大小
	staticSlotCount uint
	// 静态变量
	staticVars Slots
	// 表示类的 <clinit> 方法是否已经开始执行
	initStarted bool
	// java.lang.Class 实例
	jClass *Object
	sourceFile string
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

// 是否是 public
func (c *Class) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}

// 是否是 final
func (c *Class) IsFinal() bool {
	return 0 != c.accessFlags&ACC_FINAL
}

// 用来表示如何调用父类的方法，是否使用 invokespecial 指令，JDK 1.02 后为真
func (c *Class) IsSuper() bool {
	return 0 != c.accessFlags&ACC_SUPER
}

// 是否是接口
func (c *Class) IsInterface() bool {
	return 0 != c.accessFlags&ACC_INTERFACE
}

// 是否是抽象类
func (c *Class) IsAbstract() bool {
	return 0 != c.accessFlags&ACC_ABSTRACT
}

// 该类是否是编译器合成代码
func (c *Class) IsSynthetic() bool {
	return 0 != c.accessFlags&ACC_SYNTHETIC
}

// 是否是注解
func (c *Class) IsAnnotation() bool {
	return 0 != c.accessFlags&ACC_ANNOTATION
}

// 是否是枚举类
func (c *Class) IsEnum() bool {
	return 0 != c.accessFlags&ACC_ENUM
}

// 是否是可访问的 public 或 同一个包中，即是否有访问权限
func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() ||
		c.getPackageName() == other.getPackageName()
}

// 获取包名
func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

// 获取 main 方法
func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

// 获取初始化方法
func (c *Class) GetClinitMethod() *Method {
	return c.getStaticMethod("<clinit>", "()V")
}

// 获取包名
func (c *Class) GetPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

// 根据参数获取相关静态方法
func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range c.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

// 新建对象
func (c *Class) NewObject() *Object {
	return newObject(c)
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) Fields() []*Field {
	return c.fields
}

func (c *Class) Methods() []*Method {
	return c.methods
}

func (c *Class) JClass() *Object {
	return c.jClass
}

func (c *Class) SetJClass(jClass *Object) {
	c.jClass = jClass
}

// 转换成 Java 字符串（格式：例如 java.lang.Object）
func (c *Class) JavaName() string {
	return strings.Replace(c.name, "/", ".", -1)
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) Loader() *ClassLoader {
	return c.loader
}

// 返回与类对应的数组类
func (c *Class) ArrayClass() *Class {
	// 根据类名得到数组类名
	arrayClassName := getArrayClassName(c.name)
	// 调用类加载器加载数组类
	return c.loader.LoadClass(arrayClassName)
}

func (c *Class) isJlObject() bool {
	return c.name == "java/lang/Object"
}

func (c *Class) isJlCloneable() bool {
	return c.name == "java/lang/Cloneable"
}

func (c *Class) isJioSerializable() bool {
	return c.name == "java/io/Serializable"
}

// 判断类是否是基本类型的类
func (c *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[c.name]
	return ok
}

// 根据字段名和描述符查找字段
func (c *Class) getField(name, descriptor string, isStatic bool) *Field {
	for cl := c; cl != nil; cl = cl.superClass {
		for _, field := range cl.fields {
			if field.IsStatic() == isStatic &&
					field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (c *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for cl := c; cl != nil; cl = cl.superClass {
		for _, method := range cl.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (c *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := c.getField(fieldName, fieldDescriptor, true)
	return c.staticVars.GetRef(field.slotId)
}
func (c *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := c.getField(fieldName, fieldDescriptor, true)
	c.staticVars.SetRef(field.slotId, ref)
}

func (c *Class) GetInstanceMethod(name, descriptor string) *Method {
	return c.getMethod(name, descriptor, false)
}

func (c *Class) SourceFile() string {
	return c.sourceFile
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknown"
}