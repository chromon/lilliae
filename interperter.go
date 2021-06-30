package main

import (
	"fmt"
	"lilliae/instructions"
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// 解释器
// logInst 控制是否把指令执行信息打印到控制台
func interpret(method *heap.Method, logInst bool, args []string) {
	// 创建 Thread 实例
	thread := runtimedataarea.NewThread()
	// 创建栈帧并推入虚拟机栈顶
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	// 将 args 参数转换成 java 字符串数组
	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catchErr(thread)
	loop(thread, logInst)
}

// 将 args 参数转换成 java 字符串数组
func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}

func catchErr(thread *runtimedataarea.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

// 循环执行计算 PC、解码指令、执行指令三个步骤，直到遇到错误退出
func loop(thread *runtimedataarea.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		// 当前帧
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 根据 pc 从当前方法中解码出一条指令
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// 执行指令
		inst.Execute(frame)
		// 判断 java 虚拟机栈中是否还有栈帧，没有则退出循环
		if thread.IsStackEmpty() {
			break
		}
	}
}

// 在方法执行过程中打印指令信息
func logInstruction(frame *runtimedataarea.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

// 打印虚拟机栈信息
func logFrames(thread *runtimedataarea.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}