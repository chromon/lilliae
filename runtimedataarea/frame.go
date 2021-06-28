package runtimedataarea

import "lilliae/runtimedataarea/heap"

// 虚拟机栈栈帧（包括局部变量表、操作数栈、方法返回值和动态链接）
type Frame struct {
	// 用来实现链表数据结构指向下一个元素
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 操作数栈指针
	operandStack *OperandStack
	// 线程和下一条指令地址，为了实现跳转指令而添加的
	thread *Thread
	method *heap.Method
	nextPC int
}

// 局部变量表大小和操作数栈深度是由编译器预先计算好的，
// 存储在 class 文件 method_info 结构的 Code 属性中
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame {
		thread: thread,
		method: method,
		localVars: newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}
func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}