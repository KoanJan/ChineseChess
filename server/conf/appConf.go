package conf

import (
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// 应用配置
type appConf struct {
	Mongo  *mongoConf  `yaml:"mongoConf"`
	Redis  *redisConf  `yaml:"redisConf"`
	Logger *loggerConf `yaml:"loggerConf"`
}

// mongo配置
type mongoConf struct {
	Host     string `yaml:"host"`     // 主机
	Port     string `yaml:"port"`     // 端口
	DBName   string `yaml:"dbName"`   // 数据库
	Username string `yaml:"username"` // 用户名
	Password string `yaml:"password"` // 密码
}

// redis配置
type redisConf struct {
	Address string `yaml:"address"` // 地址
}

// 日志配置
type loggerConf struct {
	Path string `yaml:"path"` // 日志文件所在目录
}

// 获取日志路径
func (this *loggerConf) FilePath() string {
	pth := this.Path
	if !strings.HasSuffix(pth, "/") {
		pth = pth + "/"
	}
	return pth + time.Now().Format("2006-01-02") + ".log"
}

var AppConf *appConf = new(appConf)

// 初始化配置
func init() {

	data, err := ioutil.ReadFile("server/conf/appConf.yml")
	if err != nil {
		panic("读取配置文件信息出错!")
	}
	if err = yaml.Unmarshal(data, AppConf); err != nil {
		panic("读取配置文件信息出错!")
	}
}
