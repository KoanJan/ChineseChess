package conf

import (
	"strings"
	"time"
)

// 应用配置
type appConf struct {
	Mongodb *mongodbConf
	Redis   *redisConf
	Logger  *loggerConf
}

// mongo配置
type mongodbConf struct {
	URL      *string // 连接
	DBName   *string // 数据库
	Username *string // 用户名
	Password *string // 密码
}

// redis配置
type redisConf struct {
	Address *string // 地址
	Select  *int    // 数据库
}

// 日志配置
type loggerConf struct {
	Dir *string // 日志文件所在目录
}

// 获取日志路径
func (this *loggerConf) FilePath() string {
	d := *this.Dir
	if !strings.HasSuffix(d, "/") {
		d = d + "/"
	}
	return d + time.Now().Format("2006-01-02") + ".log"
}

// 配置
var AppConf *appConf = newAppConf()
