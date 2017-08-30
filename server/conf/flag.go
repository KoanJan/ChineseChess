package conf

import (
	"os"
	"strings"
)

const (
	FlagConfFilePath = "conf_path"    // 配置文件路径
	FlagLogFileDir   = "log_file_dir" // 日志文件所在目录
)

var args map[string]string

func initFlag() {

	args = make(map[string]string)

	for _, arg := range os.Args {
		if strings.Contains(arg, "=") {
			kv := strings.Split(arg, "=")
			args[kv[0]] = kv[1]
		}
	}
}

// 获取命令行参数
func FlagArgs(key string) (string, bool) {
	v, ok := args[key]
	return v, ok
}
