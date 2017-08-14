package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 应用配置
type appConf struct {
	Mongo *mongoConf `yaml:"mongoConf"`
}

// mongo配置
type mongoConf struct {
	Host     string `yaml:"host"`     // 主机
	Port     string `yaml:"port"`     // 端口
	DBName   string `yaml:"dbName"`   // 数据库
	Username string `yaml:"username"` // 用户名
	Password string `yaml:"password"` // 密码
}

var AppConf *appConf = new(appConf)

// 初始化配置
func init() {

	data, err := ioutil.ReadFile("server/conf/appConf.yml")
	if err != nil {
		fmt.Println(err)
		panic("读取配置文件信息出错!")
	}
	if err = yaml.Unmarshal(data, AppConf); err != nil {
		panic("读取配置文件信息出错!")
	}
}
