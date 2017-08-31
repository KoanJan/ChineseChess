package conf

import (
	"flag"
)

type FlagName string

const (
	FN_LogFileDir FlagName = "log_file_dir"
)

// 初始化终端参数
func initFlag() {
	flag.String(string(FN_LogFileDir), "", "[optional]日志文件所在目录[-log_file_dir=<your_log_file_directory>]")
	flag.Parse()
}

// 获取终端参数
func FlagValue(name FlagName) flag.Value {
	if f := flag.CommandLine.Lookup(string(name)); f != nil {
		return f.Value
	}
	return nil
}
