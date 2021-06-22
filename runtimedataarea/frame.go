package runtimedataarea

// 虚拟机栈栈帧（包括局部变量表、操作数栈、方法返回值和动态链接）
type Frame struct {
	// 用来实现链表数据结构指向下一个元素
	lower *Frame
	// 局部变量表指针
	LocalVars LocalVars
	// 操作数栈指针
	OperandStack *OperandStack
}

// 局部变量表大小和操作数栈深度是由编译器预先计算好的，
// 存储在 class 文件 method_info 结构的 Code 属性中
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame {
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
	}
}