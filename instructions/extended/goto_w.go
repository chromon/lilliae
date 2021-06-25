package extended

import (
	"lilliae/instructions/base"
	"lilliae/runtimedataarea"
)

// goto_w 指令和 goto 指令的唯一区别就是索引从 2 字节变成了 4 字节
type GOTO_W struct {
	offset int
}

func (gtw *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	gtw.offset = int(reader.ReadInt32())
}

func (gtw *GOTO_W) Execute(frame *runtimedataarea.Frame) {
	base.Branch(frame, gtw.offset)
}