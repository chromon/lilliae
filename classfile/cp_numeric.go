package classfile

import "math"

/*
CONSTANT_Integer_info {
	u1 tag;
	u4 bytes;
}
*/

// CONSTANT_Integer_info 使用 4 字节存储整数常量
// CONSTANT_Integer_info 正好可以容纳一个 Java 的 int 型常量，
// 但实际上比 int 更小的 boolean、byte、short 和 char 类型的常量也放在
// CONSTANT_Integer_info 中
type ConstantIntegerInfo struct {
	val int32
}

// 先读取一个 uint32 数据，然后转型为 int32 类型
func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cii.val = int32(bytes)
}

/*
CONSTANT_Float_info {
	u1 tag;
	u4 bytes;
}
 */

// CONSTANT_Float_info 使用 4 字节存储 IEEE754 单精度浮点数常量
type ConstantFloatInfo struct {
	val float32
}

// 先读取一个 uint32 数据，然后调用 math 包的 Float32frombits() 函数转换成 float32 类型
func (cfi *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cfi.val = math.Float32frombits(bytes)
}

/*
CONSTANT_Long_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
 */

// CONSTANT_Long_info使用 8 字节存储整数常量
type ConstantLongInfo struct {
	val int64
}

// 先读取 uint64 数据，然后转型为 int64 类型
func (cli *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cli.val = int64(bytes)
}

/*
CONSTANT_Double_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
 */

// CONSTANT_Double_info，使用 8 字节存储 IEEE754 双精度浮点数
type ConstantDoubleInfo struct {
	val float64
}

// 读取 uint64 数据，调用 math 包 Float64frombits() 函数转型为 float64 类型
func (cdi *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cdi.val = math.Float64frombits(bytes)
}