package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式类路径
// java -cp path/classes ...
type DirEntry struct {
	// 目录的绝对路径
	absDir string
}

// 构建实例
func newDirEntry(path string) *DirEntry {
	// 将 path 转为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 创建 DirEntry 实例并返回
	return &DirEntry{absDir}
}

// 读取 class 文件
func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 将目录和 class 文件名拼接为完整路径
	fileName := filepath.Join(de.absDir, className)
	// 读取 class 文件内容并返回
	data, err := ioutil.ReadFile(fileName)
	return data, de, err
}

func (de *DirEntry) String() string {
	return de.absDir
}