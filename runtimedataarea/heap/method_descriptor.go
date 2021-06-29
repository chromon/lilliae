package heap

// 方法描述信息
type MethodDescriptor struct {
	// 参数类型列表
	parameterTypes []string
	// 返回值类型
	returnType     string
}

func (d *MethodDescriptor) addParameterType(t string) {
	pLen := len(d.parameterTypes)
	if pLen == cap(d.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, d.parameterTypes)
		d.parameterTypes = s
	}

	d.parameterTypes = append(d.parameterTypes, t)
}