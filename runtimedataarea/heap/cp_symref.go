package heap

// 类、字段、方法和接口方法的符号引用具有一定共性，抽象为结构体
type SymRef struct {
	// 符号引用所在的运行时常量池指针
	cp *ConstantPool
	// 类的完全限定名
	className string
	// 缓存解析后的 class 结构体指针
	// 对于类符号引用只要有类名，就可以解析符号引用
	// 对于字段，需要解析类符号引用得到类数据，然后用字段名和描述符查找字段数据
	// 对于方法符号引用解析过程和字段符号引用类似
	class *Class
}

// 解析类符号引用
func (sr *SymRef) ResolvedClass() *Class {
	// 如果类符号引用已经解析直接返回
	if sr.class == nil {
		// 否则解析类符号引用
		sr.resolveClassRef()
	}
	return sr.class
}

// 根据 Java 虚拟机规范 5.4.3.1节给出了类符号引用的解析步骤
// 如果类 D 通过符号引用 N 引用类 C 的话，要解析 N，
// 先用 D 的类加载器加载 C，然后检查 D 是否有权限访问 C，
// 如果没有，则抛出 IllegalAccessError 异常
func (sr *SymRef) resolveClassRef() {
	d := sr.cp.class
	c := d.loader.LoadClass(sr.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	sr.class = c
}