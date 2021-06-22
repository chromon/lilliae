package runtimedataarea

// 虚拟机栈（使用链表实现）
type Stack struct {
	// 栈的容量
	maxSize uint
	// 栈的当前大小
	size uint
	// 栈顶指针
	_top *Frame
}

// 新建栈
func newStack(maxSize uint) *Stack {
	return &Stack {
		maxSize: maxSize,
	}
}

func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		// 当前栈帧的 lower 指向栈顶元素
		frame.lower = s._top
	}

	// 栈顶指向当前栈帧
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty")
	}

	top := s._top
	// 栈顶指针直线当前栈帧的下一个元素
	s._top = top.lower
	top.lower = nil
	s.size--
	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty")
	}
	return s._top
}