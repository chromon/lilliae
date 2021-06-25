package classfile

type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}

func (ua *UnparsedAttribute) readInfo(reader *ClassReader) {
	ua.info = reader.readBytes(ua.length)
}

func (ua *UnparsedAttribute) Info() []byte {
	return ua.info
}