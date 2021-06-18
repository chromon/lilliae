package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// 类路径项
// 类路径可以分为以下 3 个部分：
// 启动类路径（bootstrap classpath）：jre/lib，主要是 java 标准库
// 扩展类路径（extension classpath）：jre/lib/ext，主要是 java 扩展机制类
// 和用户类路径（user classpath）：默认是当前目录，可以使用 java -cp 命令指定
// 		主要包含类运行所依赖其他用户类以及第三方类库的路径，通常是类库，jar包等
type Entry interface {
	// 寻找并加载 class 文件
	// className 为 class 文件的相对路径，使用 / 分隔，文件后缀 .class
	// 例如读取 java.lang.Object 类，传入的参数为 java/lang/Object.class
	// 返回值为读取到的字节数，最终定位到 class 文件的 Entry，以及错误信息
	readClass(className string) ([]byte, Entry, error)
	// 返回字符串表示 （toString）
	String() string
}

// 根据参数创建 Entry 实例
func newEntry(path string) Entry {
	// 同时指定多个目录或文件，以分隔符分隔
	// java -cp path/classes;lib/a.jar;lib/b.zip ...
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// 使用通配符，JDK 6 开始可以使用通配符（*） 指定某个目录下所有 jar 文件
	// java -cp classes/* ...
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// zip 或 jar 文件形式类路径
	// java -cp path/lib.zip(jar) ...
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 目录形式类路径
	// java -cp path/classes ...
	return newDirEntry(path)
}

