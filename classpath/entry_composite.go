package classpath

import (
	"errors"
	"strings"
)

// 多类路径
// java -cp path/classes;lib/a.jar;lib/b.zip ...
type CompositeEntry []Entry

// 构造实例
func newCompositeEntry(pathList string) CompositeEntry {
	// warning: Entry slice declaration via literal
	// compositeEntry := []Entry{}
	// suggestion:
	var compositeEntry []Entry
	// or:
	// compositeEntry := make([]Entry, 0)

	// 将路径列表按分隔符分割成多个小路径，并将小路径转换为具体的实例
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 读取 class 文件
func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 依次调用每个子路径的 readClass() 方法，成功读取 class 数据返回
	for _, entry := range ce {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	// 如果遍历完所有子路径还没有找到 class 文件，返回错误
	return nil, nil, errors.New("class not found: " + className)
}

func (ce CompositeEntry) String() string {
	// 依次调用每个子路径的 String() 方法，拼接即可
	strs := make([]string, len(ce))
	for i, entry := range ce {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}