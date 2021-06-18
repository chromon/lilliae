package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	// 启动类路径
	bootClasspath Entry
	// 扩展类路径
	extClasspath Entry
	// 用户类路径
	userClasspath Entry
}

// 解析类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// 使用 -Xjre 选项解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	// 使用 -classpath/-cp 选项解析用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析启动类路径和扩展类路径
func (cp *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 获取 jre 目录
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户类路径
func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

// 如果没有提供 -classpath/-cp 选项，则使用当前目录作为用户路径
// ReadClass 方法依次从启动类路径、扩展类路径和用户类路径搜索 class 文件
func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

// 获取 jre 目录
// 优先使用输入的 -Xjre 选项作为 jre 目录，没有则在当前目录下寻找
// 如果还没有则尝试使用 JAVA_HOME 环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if javaHome := os.Getenv("JAVA_HOME"); javaHome != "" {
		return filepath.Join(javaHome, "jre")
	}
	panic("can not find jre folder")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false;
		}
	}
	return true
}

