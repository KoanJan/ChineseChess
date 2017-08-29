package logger

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 文件是否存在
func fileExists(filename string) bool {

	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// NewLogFileWriter returns a LogFileWriter
func NewLogFileWriter(filename string) io.WriteCloser {

	// 绝对路径
	if !filepath.IsAbs(filename) {
		filename, _ = filepath.Abs(filename)
	}

	idx := strings.LastIndex(filename, "/")
	filenameBytes := []byte(filename)

	dir := string(filenameBytes[0:idx])   // 目录
	name := string(filenameBytes[idx+1:]) // 文件名

	if name == "" {
		panic("日志文件名称错误!")
	}

	// 如果是根目录
	if dir == "" {
		dir = "/"
	}

	// 如果所在目录不存在则先创建目录
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}

	var (
		w   io.WriteCloser
		err error
	)
	// 如果日志文件不存在则先创建日志文件
	if fileExists(filename) {
		w, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	} else {
		w, err = os.Create(filename)
	}
	if err != nil {
		panic(err)
	}

	return w
}
