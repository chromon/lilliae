package heap

import "math"

// 变量槽，表示类变量和实例变量（与运行时数据区中的 slot 基本相同，防止互相引用）
type Slot struct {
	// 存放整数
	num int32
	// 存放引用
	ref *Object
}

// 类变量和实例变量表
type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

// 存放 int 类型变量
func (s Slots) SetInt(index uint, val int32) {
	s[index].num = val
}

// 获取 int 类型变量
func (s Slots) GetInt(index uint) int32 {
	return s[index].num
}

// 存放 float 类型变量
func (s Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	s[index].num = int32(bits)
}

// 获取 float 类型变量
func (s Slots) GetFloat(index uint) float32 {
	bits := uint32(s[index].num)
	return math.Float32frombits(bits)
}

// 存放 long 类型变量，占两位
func (s Slots) SetLong(index uint, val int64) {
	s[index].num = int32(val)
	s[index + 1].num = int32(val >> 32)
}

// 读取 long 类型变量
func (s Slots) GetLong(index uint) int64 {
	low := uint32(s[index].num)
	high := uint32(s[index+1].num)
	return int64(high)<<32 | int64(low)
}

// 设置 double 类型变量，占两位
func (s Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	s.SetLong(index, int64(bits))
}

// 读取 double 类型变量
func (s Slots) GetDouble(index uint) float64 {
	bits := uint64(s.GetLong(index))
	return math.Float64frombits(bits)
}

// 存放引用类型变量
func (s Slots) SetRef(index uint, ref *Object) {
	s[index].ref = ref
}

// 获取引用类型变量
func (s Slots) GetRef(index uint) *Object {
	return s[index].ref
}
