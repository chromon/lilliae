package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// 把类名转变成类型描述符，然后在前面加上方括号即可
// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// 根据数组类名推测出数组元素类名
// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	// 数组类名以方括号开头，把它去掉就是数组元素的类型描述符，
	// 然后把类型描述符转成类名即可
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// 由类名得到描述符
// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	// 如果是数组名，描述符就是类名，直接返回
	if className[0] == '[' {
		// array
		return className
	}

	// 如果是基本类型名，返回对应的类型描述符
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// 普通的类名在前面加上 [;
	return "L" + className + ";"
}

// 把类型描述符转成类名
// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor {
			// primitive
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}