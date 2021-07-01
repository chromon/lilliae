package lang

import (
	"fmt"
	"lilliae/native"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

const jlThrowable = "java/lang/Throwable"

// 用来记录 Java 虚拟机栈帧信息
type StackTraceElement struct {
	// 类所在的文件名
	fileName   string
	// 声明方法的类名
	className  string
	// 方法名
	methodName string
	// 帧正在执行哪行代码
	lineNumber int
}

func (e *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		e.className, e.methodName, e.fileName, e.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *runtimedataarea.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *runtimedataarea.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

// 计算所需跳过的帧数
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

// 根据帧创建 StackTraceElement 实例
func createStackTraceElement(frame *runtimedataarea.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}