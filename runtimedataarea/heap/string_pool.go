package heap

import "unicode/utf16"

// 在 class 文件中，字符串是以 MUTF8 格式保存
// 在 String 对象内部，字符串又是以 UTF16 格式保存的

// 使用 map 表示字符串常量池，key 是 GO 字符串， value 是 Java 字符串
var internedStrings = map[string]*Object{}

// 将 go string 转为响应 java.lang.String 实例
func JString(loader *ClassLoader, goStr string) *Object {
	// 如果 Java 字符串已经在池中，直接返回即可
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	// 先把 Go 字符串（UTF8 格式）转换成 Java 字符数组（UTF16 格式）
	chars := stringToUtf16(goStr)
	// 创建 java 字符串对象实例
	jChars := &Object{loader.LoadClass("[C"), chars, nil}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	// 将 java 字符串的 value 变量设置成刚刚转换而来的字符数组
	jStr.SetRefVar("value", "[C", jChars)

	// 将 java 字符串放入常量池
	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	// 先拿到 String 对象的 value 变量值
	charArr := jStr.GetRefVar("value", "[C")
	// 把字符数组转换成 Go 字符串
	return utf16ToString(charArr.Chars())
}

// 将 Go 字符串（UTF8 格式）转换成 Java 字符数组（UTF16 格式）
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// 把字符数组转换成 Go 字符串
// utf16 -> utf8
func utf16ToString(s []uint16) string {
	// 将 UTF16 数据转换成 UTF8 编码
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	// 强转成 go 字符串
	return string(runes)
}

// 检查字符串常量池中是否有当前字符串，如果没有则放入并返回该字符串，否则找到并直接返回
func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}