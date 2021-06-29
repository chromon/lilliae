package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (p *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	p.raw = descriptor
	p.parsed = &MethodDescriptor{}
	p.startParams()
	p.parseParamTypes()
	p.endParams()
	p.parseReturnType()
	p.finish()
	return p.parsed
}

func (p *MethodDescriptorParser) startParams() {
	if p.readUint8() != '(' {
		p.causePanic()
	}
}
func (p *MethodDescriptorParser) endParams() {
	if p.readUint8() != ')' {
		p.causePanic()
	}
}
func (p *MethodDescriptorParser) finish() {
	if p.offset != len(p.raw) {
		p.causePanic()
	}
}

func (p *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + p.raw)
}

func (p *MethodDescriptorParser) readUint8() uint8 {
	b := p.raw[p.offset]
	p.offset++
	return b
}
func (p *MethodDescriptorParser) unreadUint8() {
	p.offset--
}

func (p *MethodDescriptorParser) parseParamTypes() {
	for {
		t := p.parseFieldType()
		if t != "" {
			p.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (p *MethodDescriptorParser) parseReturnType() {
	if p.readUint8() == 'V' {
		p.parsed.returnType = "V"
		return
	}

	p.unreadUint8()
	t := p.parseFieldType()
	if t != "" {
		p.parsed.returnType = t
		return
	}

	p.causePanic()
}

func (p *MethodDescriptorParser) parseFieldType() string {
	switch p.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return p.parseObjectType()
	case '[':
		return p.parseArrayType()
	default:
		p.unreadUint8()
		return ""
	}
}

func (p *MethodDescriptorParser) parseObjectType() string {
	unread := p.raw[p.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		p.causePanic()
		return ""
	} else {
		objStart := p.offset - 1
		objEnd := p.offset + semicolonIndex + 1
		p.offset = objEnd
		descriptor := p.raw[objStart:objEnd]
		return descriptor
	}
}

func (p *MethodDescriptorParser) parseArrayType() string {
	arrStart := p.offset - 1
	p.parseFieldType()
	arrEnd := p.offset
	descriptor := p.raw[arrStart:arrEnd]
	return descriptor
}