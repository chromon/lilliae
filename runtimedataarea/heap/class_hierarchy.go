package heap

// 三种情况，S类型的引用值可以赋值给T类型：
// S 和 T 是同一类型；
// T 是类且 S 是 T 的子类；
// 或者 T 是接口且 S 实现了 T 接口
func (c *Class) isAssignableFrom(other *Class) bool {
	s, t := other, c

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// 判断是否是子类，判断 S 是否是 T 的子类，实际上也就是判断 T 是否是 S 的（直接或间接）超类
func (c *Class) IsSubClassOf(other *Class) bool {
	for c := c.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// 判断是否实现接口
// 判断 S 是否实现了 T 接口，就看 S 或 S 的（直接或间接）超类是否
// 实现了某个接口 T'，T' 要么是 T，要么是 T 的子接口
func (c *Class) IsImplements(iface *Class) bool {
	for c := c; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// c extends iface
func (c *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// class extends c
func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}