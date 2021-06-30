package heap

// 对象
type Object struct {
	// 存放对象的 Class 指针
	class *Class
	// 存放实例变量
	data interface{}
	// 用来记录 Object 结构体实例的额外信息
	extra interface{}
}

// 新建对象
func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

// getters
func (obj *Object) Class() *Class {
	return obj.class
}

func (obj *Object) Fields() Slots {
	return obj.data.(Slots)
}

// 判断对象是否是某个类的实例
func (obj *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(obj.class)
}

// 直接给对象的引用类型实例变量赋值
func (obj *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (obj *Object) GetRefVar(name, descriptor string) *Object {
	field := obj.class.getField(name, descriptor, false)
	slots := obj.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (obj *Object) Extra() interface{} {
	return obj.extra
}

func (obj *Object) SetExtra(extra interface{}) {
	obj.extra = extra
}