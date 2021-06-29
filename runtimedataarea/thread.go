package runtimedataarea

import "lilliae/runtimedataarea/heap"

// 线程
type Thread struct {
	pc int
	// 栈
	stack *Stack
}

func NewThread() *Thread {
	return &Thread {
		// newStack 方法为创建虚拟机栈示例，并指定栈帧数
		stack: newStack(1024),
	}
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

// 压栈
func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

// 弹栈
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

// 返回栈顶元素
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

// 返回栈顶元素
func (t *Thread) TopFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(t, method)
}

func (t *Thread) IsStackEmpty() bool {
	return t.stack.isEmpty()
}