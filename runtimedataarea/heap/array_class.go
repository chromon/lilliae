package heap

// 数组类特有的方法

// 判断类是否是数组
func (c *Class) IsArray() bool {
	return c.name[0] == '['
}

// 返回数组类的元素类型
func (c *Class) ComponentClass() *Class {
	// 根据数组类名推测出数组元素类名
	componentClassName := getComponentClassName(c.name)
	// 用类加载器加载元素类
	return c.loader.LoadClass(componentClassName)
}

// 创建数组对象
func (c *Class) NewArray(count uint) *Object {
	if !c.IsArray() {
		// 如果类不是数组就抛异常
		panic("Not array class: " + c.name)
	}

	// 根据数组类型创建数组对象
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count)}
	case "[B":
		return &Object{c, make([]int8, count)}
	case "[C":
		return &Object{c, make([]uint16, count)}
	case "[S":
		return &Object{c, make([]int16, count)}
	case "[I":
		return &Object{c, make([]int32, count)}
	case "[J":
		return &Object{c, make([]int64, count)}
	case "[F":
		return &Object{c, make([]float32, count)}
	case "[D":
		return &Object{c, make([]float64, count)}
	default:
		return &Object{c, make([]*Object, count)}
	}
}