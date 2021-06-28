package heap

// accessFlags 是类的访问标志，字段和方法也有访问标志，但含义可能不同
// 根据 Java 虚拟机规范，各标志定义
const (
	// 声明为 public，可以在外部访问
	ACC_PUBLIC       = 0x0001 // class field method
	// 声明为 private，尽可在类中访问
	ACC_PRIVATE      = 0x0002 //       field method
	// 声明为 protected，可以在子类中访问
	ACC_PROTECTED    = 0x0004 //       field method
	// 声明为 static
	ACC_STATIC       = 0x0008 //       field method
	// 声明为 final，不能被覆盖
	ACC_FINAL        = 0x0010 // class field method
	// 用来表示如何调用父类的方法，invokenonvirtual 和 invokespecial
	ACC_SUPER        = 0x0020 // class
	// 声明为 synchronized
	ACC_SYNCHRONIZED = 0x0020 //             method
	// 声明为 volatile
	ACC_VOLATILE     = 0x0040 //       field
	// 声明为桥接方法，由编译器生成
	ACC_BRIDGE       = 0x0040 //             method
	// 声明为 transient
	ACC_TRANSIENT    = 0x0080 //       field
	// 声明可变数量的参数
	ACC_VARARGS      = 0x0080 //             method
	// 声明为 native
	ACC_NATIVE       = 0x0100 //             method
	// 声明为借口
	ACC_INTERFACE    = 0x0200 // class
	// 声明为抽象类
	ACC_ABSTRACT     = 0x0400 // class       method
	// 声明为 strictfp，浮点模式采用 FP-strict 模式
	ACC_STRICT       = 0x0800 //             method
	// 声明为 synthetic，不在源码中
	ACC_SYNTHETIC    = 0x1000 // class field method
	// 声明为注解
	ACC_ANNOTATION   = 0x2000 // class
	// 声明为枚举
	ACC_ENUM         = 0x4000 // class field
)