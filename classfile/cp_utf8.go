package classfile

import (
	"fmt"
	"unicode/utf16"
)

// 字符串在 class 文件中是以 MUTF-8（Modified UTF-8）方式编码的

/*
CONSTANT_Utf8_info {
	u1 tag;
	u2 length;
	u1 bytes[length];
}
 */

// CONSTANT_Utf8_info 常量里放的是 MUTF-8 编码的字符串
type ConstantUtf8Info struct {
	str string
}

// 先读取除 []byte，后调用 decodeMUTF8() 函数解码成 GO 字符串
func (cui *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	cui.str = decodeMUTF8(bytes)
}

// MUTF-8 编码方式和 UTF-8 大致相同，但并不兼容。
// 差别有两点：
// 1. null 字符（代码点 U+0000）会被编码成 2 字节 0xC0、0x80
// 2. 补充字符（Supplementary Characters，代码点大于 U+FFFF 的 Unicode 字符）
// 是按 UTF-16 拆分为代理对（Surrogate Pair） 分别编码的

// 因为 GO 语言字符串使用 UTF-8 编码，所以如果字符串中不包含 null 字符或补充字符时，
// 如下简化版 readMUTF8() 也可以工作
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

// Java 序列化机制也使用了 MUTF-8 编码，java.io.DataInput 和 java.io.DataOutput 接口
// 分别定义了 readUTF() 和 WriteUTF() 方法，可以读写 MUTF-8 编码字符串
// 实际过程：mutf8 -> utf16 -> utf32 -> string
func decodeMUTF8Complete(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}