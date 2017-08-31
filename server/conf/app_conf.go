package conf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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
	Select  int    `yaml:"select"`  // 数据库
}

// 日志配置
type loggerConf struct {
	Dir string `yaml:"dir"` // 日志文件所在目录
}

// 获取日志路径
func (this *loggerConf) FilePath() string {
	d := this.Dir
	if !strings.HasSuffix(d, "/") {
		d = d + "/"
	}
	return d + time.Now().Format("2006-01-02") + ".log"
}

var AppConf *appConf = new(appConf)

// 初始化配置
func init() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	}()

	// 读取终端参数列表
	initFlag()

	// 读取配置文件
	confPath, _ := filepath.Abs("conf/app_conf.yml")
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(data, AppConf); err != nil {
		panic(err)
	}

	// 读取日志配置
	if v := FlagValue(FN_LogFileDir); v != nil {
		AppConf.Logger.Dir = v.String()
	}
}
