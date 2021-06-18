package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// zip 或 jar 文件形式类路径
// java -cp path/lib.zip(jar) ...
type ZipEntry struct {
	// zip 或 jar 文件的绝对路径
	absPath string
}

// 构造实例
func newZipEntry(path string) *ZipEntry {
	// 将 path 转为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 创建 ZipEntry 实例并返回
	return &ZipEntry{absPath}
}

// 从 zip 或 jar 文件中读取 class 文件
func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开 zip 文件
	r, err := zip.OpenReader(ze.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	// 遍历 zip 包中的文件，查找 class 文件
	for _, f := range r.File {
		if f.Name == className {
			// 找到 class 文件打开
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			// 读取内容
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, ze, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}