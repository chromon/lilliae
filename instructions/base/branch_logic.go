package base

import "lilliae/runtimedataarea"

// 比较指令中的跳转函数
func Branch(frame *runtimedataarea.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}