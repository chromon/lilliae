package classfile

import "encoding/binary"

// 构成 class 文件的基本数据单位是字节，可以把整个 class 文件当成一个字节流来处理
// 稍大一些的数据由连续多个字节构成，这些数据在 class 文件中以大端（big-endian）方式存储

// Java 虚拟机规范定义了 u1、u2 和 u4 三种数据类型来表示 1、2和 4 字节无符号整数，
// 分别对应 Go 语言的 uint8、uint16 和 uint32 类型
// 相同类型的多条数据一般按表（table）的形式存储在 class 文件中
// 表由表头和表项（item）构成，表头是u2或u4整数。假设表头是 n，后面就紧跟着n个表项数据

// 数据读取工具
type ClassReader struct {
	data []byte
}

// u1： 一个字节无符号正数
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// u2： 2 个字节无符号正数
// class 文件中以大端（big-endian）方式存储
// BigEndian 可以从 []byte 中解码多字节数据
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// u4：4 个字节无符号正数
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

// u8：读取 uint64 类型数据
func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// 读取 uint16 表，表的大小有开头的 uint16 数据指明
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUint16()
	}
	return s
}

// 读取指定数量的字节
func (cr *ClassReader) readBytes(length uint32) []byte {
	bytes := cr.data[:length]
	cr.data = cr.data[length:]
	return bytes
}