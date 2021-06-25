package constants

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// const 指令把隐含在操作码中的常量值推入操作数栈顶

// 向操作数栈中压入 null
type ACONST_NULL struct{
	base.NoOperandsInstruction
}

func (cst *ACONST_NULL) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 向操作数栈中压入 double 类型 0
type DCONST_0 struct{
	base.NoOperandsInstruction
}

func (cst *DCONST_0) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 将 double 类型 1 压入操作数栈
type DCONST_1 struct{
	base.NoOperandsInstruction
}

func (cst *DCONST_1) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// 将 float 类型 0 压入操作数栈
type FCONST_0 struct{
	base.NoOperandsInstruction
}

func (cst *FCONST_0) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// 将 float 类型 1 压入操作数栈
type FCONST_1 struct{
	base.NoOperandsInstruction
}

func (cst *FCONST_1) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// 将 float 类型 2 压入操作数栈
type FCONST_2 struct{
	base.NoOperandsInstruction
}

func (cst *FCONST_2) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 将 int 类型 -1 压入操作数栈
type ICONST_M1 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_M1) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 将 int 类型 0 压入操作数栈
type ICONST_0 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_0) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(0)
}

// 将 int 类型 1 压入操作数栈
type ICONST_1 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_1) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(1)
}

// 将 int 类型 2 压入操作数栈
type ICONST_2 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_2) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(2)
}

// 将 int 类型 3 压入操作数栈
type ICONST_3 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_3) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(3)
}

// 将 int 类型 4 压入操作数栈
type ICONST_4 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_4) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(4)
}

// 将 int 类型 5 压入操作数栈
type ICONST_5 struct{
	base.NoOperandsInstruction
}

func (cst *ICONST_5) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushInt(5)
}

// 将 long 类型 0 压入操作数栈
type LCONST_0 struct{
	base.NoOperandsInstruction
}

func (cst *LCONST_0) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushLong(0)
}

// 将 long 类型 1 压入操作数栈
type LCONST_1 struct{
	base.NoOperandsInstruction
}

func (cst *LCONST_1) Execute(frame *runtimedataarea.Frame) {
	frame.OperandStack().PushLong(1)
}
