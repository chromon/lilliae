package heap

// 对象
type Object struct {
	// 存放对象的 Class 指针
	class *Class
	// 存放实例变量
	fields Slots
}

// 新建对象
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (obj *Object) Class() *Class {
	return obj.class
}
func (obj *Object) Fields() Slots {
	return obj.fields
}

// 判断对象是否是某个类的实例
func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}