package extended

import (
	"lilliae/instructions/base"
	"lilliae/instructions/loads"
	"lilliae/instructions/math"
	"lilliae/instructions/stores"
	"lilliae/runtimedataarea"
)

// 加载类指令、存储类指令、ret 指令和 iinc 指令需要按索引访问局部变量表，
// 索引以 uint8 的形式存在字节码中。对于大部分方法来说，局部变量表大小都不会超过256，
// 所以可以用一字节来表示索引
// 但如果有方法的局部变量表超过这一限制，需要使用 wide 指令扩展命令

// 扩展局部变量表，wide 指令改变其他指令的行为
type WIDE struct{
	// 存放被改变的指令
	modifiedInstruction base.Instruction
}

func (w *WIDE) FetchOperands(reader *base.BytecodeReader) {
	// 从字节码中读取一字节的操作码
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		// 创建子指令实例
		inst := &loads.ILOAD{}
		// 读取子指令的操作数
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		w.modifiedInstruction = inst
	case 0xa9:
		// ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (w *WIDE) Execute(frame *runtimedataarea.Frame) {
	// wide 指令只是增加了索引宽度，并不改变子指令操作，
	// 所以其 Execute 方法只要调用子指令的 Execute 方法即可
	w.modifiedInstruction.Execute(frame)
}