package main

import (
	"fmt"
	"lilliae/instructions"
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
	"lilliae/runtimedataarea/heap"
)

// 解释器
func interpret(method *heap.Method) {
	// 创建 Thread 实例
	thread := runtimedataarea.NewThread()
	// 创建栈帧并推入虚拟机栈顶
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *runtimedataarea.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("Local vars: %v\n", frame.LocalVars())
		fmt.Printf("Operand stack: %v\n", frame.OperandStack())
		panic(r)
	}
}

// 循环执行计算 PC、解码指令、执行指令三个步骤，直到遇到错误退出
func loop(thread *runtimedataarea.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		// 计算 PC
		pc := frame.NextPC()
		thread.SetPC(pc)
		// 解码指令
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// 执行指令
		fmt.Printf("PC: %2d inst: %T %v \n", pc, inst, inst)
		inst.Execute(frame)
	}
}