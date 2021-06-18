package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行选项和参数
type Cmd struct {
	helpFlag bool
	versionFlag bool
	// classpath 路径
	cpOption string
	// 类名
	class string
	// 参数
	args []string
	// 指定 jre 目录位置
	XjreOption string
}

// 使用 flag 包处理命令行选项
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	// java -? / -help
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	// java -version
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	// java -cp / -classpath 指定用户类路径
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	// 自定义非标准选项 -Xjre，用于指定 jre 目录位置
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

// 打印命令用法
func printUsage() {
	fmt.Printf("Usage: %s [-optins] class [args...]\n", os.Args[0])
}