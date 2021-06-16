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
	cpOption string
	class string
	args []string
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