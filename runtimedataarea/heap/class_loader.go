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
	return &ClassLoader{
		cp: cp,
		verboseFlag: verboseFlag,
		classMap: make(map[string]*Class),
	}
}

// 将类数据加载到方法区
func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		// 类已经加载
		return class
	}

	// 如果要加载的类是数组类，则调用 loadArrayClass 方法
	if name[0] == '[' {
		return cl.loadArrayClass(name)
	}

	return cl.loadNonArrayClass(name)
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
			panic("todo")
		}
	}
}