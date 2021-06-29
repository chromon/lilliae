package heap

// 分别针对引用类型数组和 7 种基本类型数组返回具体的数组数据

func (obj *Object) Bytes() []int8 {
	return obj.data.([]int8)
}

func (obj *Object) Shorts() []int16 {
	return obj.data.([]int16)
}

func (obj *Object) Ints() []int32 {
	return obj.data.([]int32)
}

func (obj *Object) Longs() []int64 {
	return obj.data.([]int64)
}

func (obj *Object) Chars() []uint16 {
	return obj.data.([]uint16)
}

func (obj *Object) Floats() []float32 {
	return obj.data.([]float32)
}

func (obj *Object) Doubles() []float64 {
	return obj.data.([]float64)
}

func (obj *Object) Refs() []*Object {
	return obj.data.([]*Object)
}

// 返回数组长度
func (obj *Object) ArrayLength() int32 {
	switch obj.data.(type) {
	case []int8:
		return int32(len(obj.data.([]int8)))
	case []int16:
		return int32(len(obj.data.([]int16)))
	case []int32:
		return int32(len(obj.data.([]int32)))
	case []int64:
		return int32(len(obj.data.([]int64)))
	case []uint16:
		return int32(len(obj.data.([]uint16)))
	case []float32:
		return int32(len(obj.data.([]float32)))
	case []float64:
		return int32(len(obj.data.([]float64)))
	case []*Object:
		return int32(len(obj.data.([]*Object)))
	default:
		panic("Not array!")
	}
}