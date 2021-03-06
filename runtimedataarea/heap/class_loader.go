package heap

import (
	"fmt"
	"lilliae/classfile"
	"lilliae/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/

// 类加载器
type ClassLoader struct {
	// Classpath 指针，用来搜索和读取 class 文件
	cp *classpath.Classpath
	// 记录已经加载的类数据，key 是类的完全限定名
	// 方法区是一个抽象的概念，classMap 可以看做是方法区的具体实现
	classMap map[string]*Class
	// 是否打印类的加载信息
	verboseFlag bool
}

// 创建 ClassLoader 实例
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}

	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

// 加载 java.lang.Class 类
func (cl *ClassLoader) loadBasicClasses() {
	// 加载 java.lang.Class 类
	jlClassClass := cl.LoadClass("java/lang/Class")
	// 同时又触发 java.lang.Object 等类和接口的加载
	for _, class := range cl.classMap {
		if class.jClass == nil {
			// 给已经加载的每一个类关联类对象
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

// 加载 void 和基本类型的类
func (cl *ClassLoader) loadPrimitiveClasses() {
	for primitiveType, _ := range primitiveTypes {
		// primitiveType 是 void、int、float 等类型
		cl.loadPrimitiveClass(primitiveType)
	}
}

// 生成 void 和基本类型类
// 有三点需要注意：
// 第一，void和基本类型的类名就是 void、int、float 等。
// 第二，基本类型的类没有超类，也没有实现任何接口。
// 第三，非基本类型的类对象是通过 ldc 指令加载到操作数栈中的。而基本类型的类对象，
// 虽然在 Java 代码中看起来是通过字面量获取的，但是编译之后的指令并不是 ldc，
// 而是 getstatic。每个基本类型都有一个包装类，包装类中有一个静态常量，叫作 TYPE，其中存放的就是基本类型的类
func (cl *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class {
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      cl,
		initStarted: true,
	}
	class.jClass = cl.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	cl.classMap[className] = class
}

// 将类数据加载到方法区
func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		// 类已经被加载
		return class
	}

	var class *Class
	if name[0] == '[' { // 数组类
		class = cl.loadArrayClass(name)
	} else {
		class = cl.loadNonArrayClass(name)
	}

	// 在类加载完之后，查看 java.lang.Class 是否已经加载
	if jlClassClass, ok := cl.classMap["java/lang/Class"]; ok {
		// 给类关联类对象
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}

	return class
}

// 加载非数组类
// 数组类与普通类有很大不同，数组类数据并不是来自 class 文件，
// 而是由 Java 虚拟机在运行期间生成
func (cl *ClassLoader) loadNonArrayClass(name string) *Class {
	// 找到 class 文件并把数据读取到内存
	data, entry := cl.readClass(name)
	// 解析 class 文件，生成虚拟机可用的类数据
	class := cl.defineClass(data)
	// 链接
	link(class)
	if cl.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

// 加载数组类
func (cl *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        name,
		loader:      cl,
		initStarted: true, // 数组类不需要初始化
		superClass:  cl.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			cl.LoadClass("java/lang/Cloneable"),
			cl.LoadClass("java/io/Serializable"),
		},
	}
	cl.classMap[name] = class
	return class
}

// 查找并读取 class 文件
func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 解析 class 文件，生成虚拟机可用的类数据
func (cl *ClassLoader) defineClass(data []byte) *Class {
	// 将 class 文件数据转化成 Class 结构体
	class := parseClass(data)
	class.loader = cl
	// Class 结构体的 superClass 字段存放父类名，是符号引用
	// 解析父类符号引用
	resolveSuperClass(class)
	// Class结构体的 interfaces 字段存放直接接口表，是符号引用
	// 解析接口表符号引用
	resolveInterfaces(class)
	// 将生成的 class 对象添加到方法区
	cl.classMap[class.name] = class
	return class
}

// 将 class 文件数据转化成 Class 结构体
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// 解析父类符号引用
func resolveSuperClass(class *Class) {
	// 除 java.lang.Object 以外，所有的类都有且仅有一个父类
	if class.name != "java/lang/Object" {
		// 非 Object 类需要递归调用 LoadClass 加载父类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 解析直接接口表符号引用
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			// 递归调用 LoadClass 方法加载类的每一个直接接口
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接：分为验证和准备两个阶段
func link(class *Class) {
	// 验证
	verify(class)
	// 准备
	prepare(class)
}

// 验证阶段
func verify(class *Class) {
	// todo
}

// 准备阶段
func prepare(class *Class) {
	// 计算实例字段个数，并编号
	calcInstanceFieldSlotIds(class)
	// 计算静态字段的个数，并编号
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 计算实例字段个数，并编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 给类变量分配空间，并赋初值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 如果静态变量属于基本类型或 String 类型，有 final 修饰符，
// 且它的值在编译期已知，则该值存储在 class 文件常量池中
// 从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// 字符串类型静态常量的初始化逻辑
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}