package base

// 字节码读取
type BytecodeReader struct {
	// 存放字节码
	code []byte
	// 记录读取到哪个字节
	pc int
}

// 重置
func (br *BytecodeReader) Reset(code []byte, pc int) {
	br.code = code
	br.pc = pc
}

// 读取一个字节
func (br *BytecodeReader) ReadUint8() uint8 {
	i := br.code[br.pc]
	br.pc++
	return i
}

func (br *BytecodeReader) ReadInt8() int8 {
	return int8(br.ReadUint8())
}

// 读取两个字节
func (br *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(br.ReadUint8())
	byte2 := uint16(br.ReadUint8())
	return (byte1 << 8) | byte2
}

func (br *BytecodeReader) ReadInt16() int16 {
	return int16(br.ReadUint16())
}

// 读取 4 个字节
func (br *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(br.ReadUint8())
	byte2 := int32(br.ReadUint8())
	byte3 := int32(br.ReadUint8())
	byte4 := int32(br.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// 多次读取 4 个字节
func (br *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = br.ReadInt32()
	}
	return ints
}

// tableswitch 指令操作码的后面有 0 ~ 3 字节的 padding，
// 以保证 defaultOffset 在字节码中的地址是 4 的倍数
func (br *BytecodeReader) SkipPadding() {
	for br.pc % 4 != 0 {
		br.ReadUint8()
	}
}

func (br *BytecodeReader) PC() int {
	return br.pc
}