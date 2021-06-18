package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry 实际上是一个 CompositeEntry
func newWildcardEntry(path string) CompositeEntry {
	// 删除路径末尾 *
	baseDir := path[:len(path) - 1]
	var compositeEntry []Entry

	// Walk 函数对每一个文件/目录都会调用 WalkFunc 函数类型值
	// 调用时 path 参数会包含 Walk 的 root 参数作为前缀
	// 就是说，如果 Walk 函数的 root 为 "dir"，该目录下有文件 "a"，
	// 将会使用 "dir/a" 调用 walkFn 参数
	// walkFn 参数被调用时的 info 参数是 path 指定的地址（文件/目录）的文件信息
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			// 用作 WalkFunc 类型的返回值，表示该次调用的 path 参数指定的目录应被跳过
			// 本错误不应被任何其他函数返回
			// 通配符类路径不能递归匹配子目录下的 jar 文件
			return filepath.SkipDir
		}
		// 根据后缀名选出 jar 文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	// func Walk(root string, walkFn WalkFunc) error
	// Walk 函数会遍历 root 指定的目录下的文件树，对每一个该文件树中的目录和文件
	// 都会调用 walkFn，包括 root 自身
	// 遍历 baseDir 创建 ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}