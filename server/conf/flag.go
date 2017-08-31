package conf

import (
	"flag"
)

// 终端参数名称
const (
	// Mongodb
	FN_MongodbURL      = "mongo_url"
	FN_MongodbDBName   = "mongo_db_name"
	FN_MongodbUsername = "mongo_user"
	FN_MongodbPassword = "mongo_pwd"

	// Redis
	FN_RedisAddress = "redis_addr"
	FN_RedisSelect  = "redis_select"

	// Logger
	FN_LogFileDir = "log_file_dir"
)

// 从终端参数读取配置
func newAppConf() *appConf {

	_appConf := &appConf{new(mongodbConf), new(redisConf), new(loggerConf)}

	_appConf.Mongodb.URL = flag.String(FN_MongodbURL, "127.0.0.1:27017", "Mongodb地址")
	_appConf.Mongodb.DBName = flag.String(FN_MongodbDBName, "chinese_chess", "Mongodb数据库名称")
	_appConf.Mongodb.Username = flag.String(FN_MongodbUsername, "", "Mongodb用户名")
	_appConf.Mongodb.Password = flag.String(FN_MongodbPassword, "", "Mongodb密码")

	_appConf.Redis.Address = flag.String(FN_RedisAddress, "127.0.0.1:6379", "Redis地址")
	_appConf.Redis.Select = flag.Int(FN_RedisSelect, 1, "Redis数据库")

	_appConf.Logger.Dir = flag.String(FN_LogFileDir, "/var/log/yi", "日志文件所在目录")

	flag.Parse()

	return _appConf
}
