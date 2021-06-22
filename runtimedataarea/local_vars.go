package runtimedataarea

import "math"

// 局部变量表
type LocalVars []Slot

// 创建局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// 存放 int 型数据
func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

// 读取 int 型数据（boolean、byte、short 和 char 类型可以转为 int 类型存取）
func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}

// 存放 float 型数据（先转成 int 类型，然后再按 int 变量处理）
func (lv LocalVars) SetFloat(index uint, val float32) {
	// Float32bits 方法返回浮点数 val 的二进制格式对应的 4 字节无符号正数
	bits := math.Float32bits(val)
	lv[index].num = int32(bits)
}

// 获取 float 型数据
func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	// 函数返回无符号整数 bits 对应二进制表示的 4 字节浮点数
	return math.Float32frombits(bits)
}

// 存放 long 类型数据（拆分成两个 int 变量）
func (lv LocalVars) SetLong(index uint, val int64) {
	lv[index].num = int32(val)
	lv[index + 1].num = int32(val >> 32)
}

// 获取 long 类型数据
func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index + 1].num)
	return int64(high) << 32 | int64(low)
}

// 存放 double 类型数据（先转成 long 类型，如何按照 long 类型处理）
func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

// 获取 double 类型数据
func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

// 获取引用类型数据
func (lv LocalVars) SetRef(index uint, ref *Object) {
	lv[index].ref = ref
}

// 设置引用类型数据
func (lv LocalVars) GetRef(index uint) *Object {
	return lv[index].ref
}