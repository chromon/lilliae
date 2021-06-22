package runtimedataarea

import "math"

// 操作数栈
type OperandStack struct {
	// 记录栈顶位置
	size uint
	// 操作数变量槽
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// 压入 int 型变量
func (ops *OperandStack) PushInt(val int32) {
	ops.slots[ops.size].num = val
	ops.size++
}

// 弹出 int 型变量
func (ops *OperandStack) PopInt() int32 {
	ops.size--
	return ops.slots[ops.size].num
}

// 压入 float 类型变量
func (ops *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	ops.slots[ops.size].num = int32(bits)
	ops.size++
}

// 弹出 float 类型变量
func (ops *OperandStack) PopFloat() float32 {
	ops.size--
	bits := uint32(ops.slots[ops.size].num)
	return math.Float32frombits(bits)
}

// 将 long 类型推入栈顶
func (ops *OperandStack) PushLong(val int64) {
	ops.slots[ops.size].num = int32(val)
	ops.slots[ops.size + 1].num = int32(val >> 32)
	ops.size += 2
}

// 弹出栈顶 long 型数据
func (ops *OperandStack) PopLong() int64 {
	ops.size -= 2
	low := uint32(ops.slots[ops.size].num)
	high := uint32(ops.slots[ops.size + 1].num)
	return int64(high) << 32 | int64(low)
}

// 将 double 类型推入栈顶
func (ops *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	ops.PushLong(int64(bits))
}

// 弹出 double 类型数据
func (ops *OperandStack) PopDouble() float64 {
	bits := uint64(ops.PopLong())
	return math.Float64frombits(bits)
}

// 引用类型压入栈顶
func (ops *OperandStack) PushRef(ref *Object) {
	ops.slots[ops.size].ref = ref
	ops.size++
}

// 弹出引用类型数据
func (ops *OperandStack) PopRef() *Object {
	ops.size--
	ref := ops.slots[ops.size].ref
	// 帮助 GC 回收 ref
	ops.slots[ops.size].ref = nil
	return ref
}